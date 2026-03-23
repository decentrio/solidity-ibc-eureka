# contracts/utils/

On-chain utilities for IBC protocol: protobuf encoding, header hashing, ZK verification, token management, and access control.

## Key Files

- **Encode.sol** — Protobuf wire-format encoding matching CometBFT `proto.Marshal()`. Includes `cdcEncode*` wrappers for header hashing.
- **Header.sol** — Tendermint header hash (`hashHeader`, 14 fields) and validator set hash (`hashValSet`) using Merkle tree with 1-byte prefixes.
- **WrapperVerifier.sol** — Ed25519 point decompression + SHA512 computation for Groth16 public inputs.
- **Groth16Verifier.sol** — Auto-generated from circuit VK. Do NOT manually edit.
- **IBCERC20.sol** — ERC20 wrapper for IBC-bridged tokens (Beacon Proxy).
- **Escrow.sol** — Token custody during cross-chain transfer (Beacon Proxy).
- **IBCRolesLib.sol** — Access control role definitions.
- **Predicates.sol** — Header/validator set verification logic.

## Conventions

- Any change to `Encode.sol` must pass `forge test --match-contract EncodeTest` (34 tests cross-validated against Go `proto.Marshal()`).
- Proto3 rule: skip zero-value fields.
- Merkle prefix: `bytes1(0x00)` / `bytes1(0x01)` — never use `[0x00]`.

See [docs/ARCHITECTURE.md](../../docs/ARCHITECTURE.md) for encoding pipeline details.
