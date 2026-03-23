# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a production Solidity implementation of **IBC v2 (Inter-Blockchain Communication)** for Ethereum ↔ Cosmos interoperability. It consists of three language layers:

- **Solidity** (`contracts/`): Core IBC protocol, ICS20 token transfers, Groth16 light client
- **Go** (`operator/`): Operator/relayer with ecip-gnark Groth16 prover for Ed25519 signature verification
- **Rust** (`packages/`, `programs/`): SP1 RISC-V proving programs (legacy), relayer, CosmWasm Ethereum light client
- **Go** (`e2e/`): End-to-end tests using the interchaintest framework

## Prerequisites

- **bun** for Solidity dependencies (`bun install`, not npm/yarn — npm fails on git deps like `@uniswap/permit2`)
- **just** task runner for build/test/lint commands
- **Foundry** (forge, cast) for Solidity compilation and testing
- **Go 1.25+** for the operator module
- **Docker Desktop** + **Kurtosis** for e2e testing
- **Nix** development shell available (`nix develop`) as alternative to manual tool installation

## Commands

All commands use the **just** task runner (`just` is required):

```bash
# Build
just build-contracts          # Compile Solidity contracts
just build-relayer            # Build Rust relayer (release)
just build-operator           # Build SP1 fixture operator (release)

# Test
just test-foundry             # All Solidity tests
just test-foundry testname    # Single Solidity test (e.g., testSendTransfer)
just test-cargo               # All Rust unit tests
just test-cargo testname      # Single Rust test
just test-benchmark testname  # Gas benchmark tests
just test-e2e testname        # Full e2e test by name

# Lint
just lint                     # All linters
just lint-solidity            # forge fmt + solhint + natlint
just lint-go                  # golangci-lint
just lint-rust                # cargo fmt + cargo clippy
just lint-buf                 # Protobuf linting

# Security
just slither                  # Static analysis (Slither)

# Generate
just generate-abi             # Extract ABIs from compiled contracts
just generate-fixtures-solidity  # Regenerate Solidity test fixtures

# Install
just install-operator         # Install operator binary
just install-relayer          # Install relayer binary

# E2E test suites (convenience)
just test-e2e-eureka          # IBC Eureka suite only
just test-e2e-relayer         # Relayer suite only
just test-e2e-cosmos-relayer  # Cosmos relayer suite only
```

To run a single Foundry test:
```bash
forge test --match-test testSendTransfer -vvv
```

Go operator tests:
```bash
cd operator && go test ./...              # All operator tests
cd operator && go test -race ./...        # With race detector
cd operator && go test -v ./prover/...    # Single package verbose
cd operator && go test -run TestParseTrustThreshold ./client/...  # Single test
```

Local e2e setup (Kurtosis ETH + Gaia):
```bash
bash run_node.sh
```

## Architecture

### Contract Hierarchy

```
AccessManager (OpenZeppelin RBAC)
    ↓
ICS26Router (UUPS) ← Main entry point for all IBC messages
    ├─ ICS20Transfer (UUPS) ← ICS-20 fungible token transfers
    │   ├─ IBCERC20 (Beacon Proxy) ← Wrapper for IBC-bridged tokens
    │   └─ Escrow (Beacon Proxy)   ← Token custody during transfer
    └─ SP1ICS07Tendermint (UUPS)   ← ZK light client for Cosmos chains
        └─ WrapperVerifier → Groth16Verifier (ecip-gnark Ed25519 circuit)
```

**ICS26Router** handles: packet sequencing, replay protection, timeout enforcement, and routing packets to the appropriate IBC application (e.g., ICS20Transfer).

**ICS20Transfer** implements: token locking/minting on send, token unlocking/burning on receive, and proof verification via the router.

**SP1ICS07Tendermint** verifies Cosmos chain state using Groth16 ZK proofs for Ed25519 signature verification (via ecip-gnark PreHashCircuit with 24 public inputs).

### Go Operator (`operator/`)

Bidirectional relayer (Cosmos↔Ethereum) with ecip-gnark Groth16 prover:
- `prover/`: Ed25519 signature → Groth16 proof generation (PreHashCircuit, 24 public inputs)
- `prover/extractor.go`: `ExtractValidatorSignature()` — extracts first non-absent Ed25519 signature from a LightBlock for Groth16 proving (single validator only, TODO: multi-validator)
- `prover/cmd/`: Circuit setup tool — compiles circuit, exports R1CS/PK/VK/Groth16Verifier.sol
- `services/`: Main relayer loop, UpdateCosmosClient, UpdateEthClient, batch processing (`BatchBuilder` with mutex-guarded packet queue)
- `subscriber/`: Event listeners — `SubscribeCosmos` (CometBFT WebSocket) and `SubscribeEth` (ICS26Router contract events)
- `client/`: Tendermint RPC (`tendermint.go`) + Ethereum Beacon API (`ethereum.go`) clients
- `transaction/`: Ethereum tx builder and sender
- `bindings/`: Go bindings for Solidity contracts

**ZK Proof Flow**: Ed25519 sig → decompress points → SHA512(R||A||msg) off-chain → gnark witness → `groth16.Prove()` → `ProofToBigInts()` → Solidity `WrapperVerifier.verifyProof()` → on-chain SHA512 + `Groth16Verifier.verifyProof()`

**Setup**: `cd operator && go run ./prover/cmd/ <output_dir>` — generates `r1cs.bin`, `pk.bin`, `vk.bin`, `Groth16Verifier.sol`

**Fixture generation**: `cd operator && go run ./prover/cmd/fixture/ [bin_dir] [output_path]` — generates Groth16 proof fixture JSON for Solidity tests

**Tests**: `cd operator && go test ./...` (all packages) or `cd operator && go test -race ./...` (with race detector). Unit tests exist for: `prover/`, `client/`, `keys/`, `utils/`, `services/`, `subscriber/`.

**Go module dependencies**: `operator/go.mod` uses `replace` directives for local paths to `ecip-gnark` (`../../ecip-gnark`) and `gnark` (`../../decentrio-gnark`). These point outside the repo and must be adjusted per developer's local setup.

**Operator tools** (`operator/cmd/`):
- `encode_debug/`: Outputs Go `proto.Marshal()` reference hex for all CometBFT types (Version, SimpleValidator, BlockID, PartSetHeader, Header). Used as source of truth to cross-validate Solidity's `Encode.sol` output via `test/solidity-ibc/EncodeTest.t.sol`.

**Config**: `operator/config.example.json` defines two modules: `cosmos_to_eth` (contract addresses for ICS26, WrapperVerifier, Membership, UpdateClient, Misbehaviour) and `eth_to_cosmos` (Beacon API URL, signer address).

### On-chain Encoding & Hashing (`contracts/utils/`)

- `Encode.sol`: Manual protobuf wire-format encoding for Tendermint types, matching Go `proto.Marshal()` exactly. Includes:
  - `encodeVarint`, `encodeString`: Standard protobuf primitives
  - `encodeValidator`: Encodes `SimpleValidator` with `PublicKey{Ed25519: ...}` oneof wrapper (matching `cmttypes.SimpleValidator.Marshal()`)
  - `encodeVersion`: Skips zero-value fields per proto3 spec (matching `cmtversion.Consensus.Marshal()`)
  - `encodeBlockId`, `encodePartSetHeader`: Direct proto.Marshal equivalents
  - `cdcEncodeString`, `cdcEncodeInt64`, `cdcEncodeBytes`, `cdcEncodeBytes32`: Wrap values in gogoproto wrapper types (`StringValue`, `Int64Value`, `BytesValue`) matching CometBFT's internal `cdcEncode()` used in `Header.Hash()`
  - `encodeTimestamp`: Encodes unix seconds as `google.protobuf.Timestamp` (matching `gogotypes.StdTimeMarshal()`)
- `Header.sol`: Computes Tendermint header hash (`hashHeader`) and validator set hash (`hashValSet`) using `Encode.sol` + Merkle tree. Called by `UpdateClient.sol` to verify header authenticity.
  - `hashHeader`: Always hashes 14 fields (matching CometBFT `Header.Hash()`), uses `cdcEncode*` wrappers
  - `hashValSet`: Merkle root of proto-encoded validators
  - `merkleHash`: Uses 1-byte leaf/inner prefixes (`bytes1(0x00)`/`bytes1(0x01)`) per Tendermint spec (RFC 6962)
- `HeightCmp.sol`: Height comparison utilities for IBC height tuples.

**Cross-validation**: `cd operator && go run ./cmd/encode_debug/` outputs Go `proto.Marshal()` reference hex. `forge test --match-contract EncodeTest` validates Solidity output matches Go exactly (34 tests).

### Contract Programs (`contracts/programs/`)

On-chain verification logic called by `SP1ICS07Tendermint`:
- `UpdateClient.sol`: Validates Tendermint header updates (height, time, validator set transitions)
- `Membership.sol`: Verifies ICS-23 Merkle proofs for packet commitments
- `Misbehaviour.sol`: Detects and handles validator equivocation

### SP1 Programs (RISC-V, in `programs/sp1-programs/`) [Legacy]

These run inside the SP1 proving system to generate verifiable proofs:
- `update-client`: Verify CometBFT headers, update client state
- `membership`: Verify IBC membership/non-membership proofs
- `uc-and-membership`: Combined update + membership (gas optimization)
- `misbehaviour`: Detect validator equivocation

### Rust Packages (`packages/`)

- `relayer/`: Multi-chain relayer with eth-to-cosmos, cosmos-to-eth, cosmos-to-cosmos modules
- `tendermint-light-client/`: Tendermint client types and proof verifiers
- `ethereum/`: Ethereum light client for CosmWasm
- `sp1-ics07-tendermint-prover/`: Utilities for generating SP1 proofs
- `go-abigen/`: Go bindings generated from contract ABIs

### Test Fixtures (`test/solidity-ibc/fixtures/`)

Pre-generated SP1 proof fixtures (Groth16 and PLONK) for unit testing without running the full prover. Both single-packet and aggregated-packet fixtures are included. Regenerate with `just generate-fixtures-solidity`.

### E2E Tests (`e2e/interchaintestv8/`)

Go tests using interchaintest that spin up real Ethereum (via Kurtosis) and Tendermint nodes. Requires Docker Desktop, relayer/operator binaries, and an SP1 network key. Test suites:
- `TestWithIbcEurekaTestSuite`: Full IBC transfer flows
- `TestWithRelayerTestSuite`: Relayer-specific tests
- `TestWithCosmosRelayerTestSuite`: Cosmos-to-Cosmos relay tests
- `TestWithSP1ICS07TendermintTestSuite`: Light client tests
- `TestWithMultichainTestSuite`: Multi-chain transfers

## Key Design Patterns

### Upgrade Strategy

All major contracts use UUPS proxy pattern. `IBCERC20` and `Escrow` use Beacon Proxy so ICS20Transfer can upgrade all instances atomically. See `UPGRADEABILITY.md` for the full governance process.

### Access Control Roles

Defined in `contracts/utils/IBCRolesLib.sol`:
- `RELAYER_ROLE` (1): Submit IBC packets
- `PAUSER_ROLE` (2) / `UNPAUSER_ROLE` (3): Emergency pause
- `RATE_LIMITER_ROLE` (5): Set transfer rate limits
- `DELEGATE_SENDER_ROLE` (4): Call `sendTransferWithSender` on behalf of others

### Proof Aggregation

ICS20Transfer supports batching multiple packets into a single proof submission, reducing gas cost by ~90% per packet for large batches (25–50 packets). Relevant in `ICS20Transfer.sol` `multiRecvPacket`/`multiAckPacket` methods.

## Linting & Static Analysis

- **solhint** (`.solhint.json`): Max code complexity 8, function max lines 70 (warning)
- **natlint**: NatSpec documentation linter for Solidity
- **Slither** (`.slither.config.json`): Security static analysis, excludes low/informational and dependencies
- **buf**: Protobuf linting and code generation (`buf.yaml`, `buf.gen.yaml`)

## Foundry Configuration

- Solidity version: `0.8.28`, EVM: `cancun`
- Optimizer: enabled with 10,000 runs + IR (`--via-ir`)
- Fuzz runs: 100,000 locally, 5,000 in CI
- Fixed block timestamp for test reproducibility
- Formatting: line length 120, tab width 4, double quotes

## Groth16 Verification (ecip-gnark)

**Circuit**: `PreHashCircuit` verifies Ed25519 signature with pre-computed hash. 24 public inputs = 6 field elements (R.X, R.Y, S, Hash, A.X, A.Y) x 4 limbs (64-bit LE).

**On-chain contracts**:
- `WrapperVerifier.sol`: Decompresses Ed25519 points, computes SHA512(R||A||msg), reduces mod L, encodes 24 limbs, calls Groth16Verifier
- `Groth16Verifier.sol`: Auto-generated from circuit VK via `operator/prover/cmd/`
- `IVerifier.sol`: Interface — `IVerifier` (wrapper) and `IGroth16Verifier` (raw)

**MsgUpdateClient** includes: `proof[8]`, `commitments[2]`, `commitmentPok[2]`, `signature[2]` (R,S as bytes32), `validatorPubkey` (bytes32), `voteSignBytes` (bytes32, keccak256 hash)

## Regenerating Go Bindings

After modifying Solidity contracts that the operator depends on (SP1ICS07Tendermint, ICS26Router, UpdateClient, Misbehaviour):
```bash
bun install && forge build
# Then use abigen to regenerate from out/<Contract>.sol/<Contract>.json
abigen --abi <abi_json> --bin <bin_hex> --pkg <PkgName> --out operator/bindings/<PkgName>/binding.go
```

ABI and Bin bytecode must stay in sync — if you update the ABI (e.g., add fields to a struct), you must also recompile and update the Bin.

## Shadowfork Tests

Tests in `test/shadowfork/` fork mainnet/testnet and require `ETH_RPC_URL` environment variable. Run separately from standard unit tests.
