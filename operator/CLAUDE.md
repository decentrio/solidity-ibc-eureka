# operator/

Go bidirectional relayer (Cosmos ↔ Ethereum) with ecip-gnark Groth16 prover for Ed25519 signature verification.

## Key Directories

- **prover/** — Ed25519 → Groth16 proof generation (PreHashCircuit, 24 public inputs)
- **prover/cmd/** — Circuit setup tool (R1CS/PK/VK/Groth16Verifier.sol export)
- **client/** — Tendermint RPC + Ethereum Beacon API clients
- **services/** — Main relay loop, UpdateCosmosClient, UpdateEthClient, BatchBuilder
- **subscriber/** — CometBFT WebSocket + Ethereum contract event listeners
- **transaction/** — Ethereum tx builder and sender
- **bindings/** — Auto-generated Go bindings for Solidity contracts
- **cmd/encode_debug/** — Outputs Go `proto.Marshal()` reference hex for cross-validating Encode.sol

## Conventions

- `go.mod` uses `replace` directives for local `ecip-gnark` and `gnark` — adjust per dev setup
- After modifying Solidity contracts: regenerate bindings with `abigen`
- Tests: `go test ./...` or `go test -race ./...`

## Config

See `config.example.json` for module configuration (contract addresses, RPC endpoints).

See [docs/ARCHITECTURE.md](../docs/ARCHITECTURE.md) for proof flow details.
