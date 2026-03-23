# e2e/interchaintestv8/

Go end-to-end tests using the interchaintest framework. Spins up real Ethereum (via Kurtosis) and CometBFT nodes.

## Test Suites

- **ibc_eureka_test.go** — Full ICS20 transfer flows
- **relayer_test.go** — Relayer-specific tests
- **cosmos_relayer_test.go** — Cosmos-to-Cosmos relay
- **sp1_ics07_test.go** — Light client verification
- **multichain_test.go** — Multi-chain transfers

## Prerequisites

Docker Desktop, Kurtosis, compiled operator/relayer binaries, SP1 network key.

## Running

```bash
just test-e2e-eureka          # IBC Eureka suite
just test-e2e-relayer         # Relayer suite
just test-e2e-cosmos-relayer  # Cosmos relayer suite
just test-e2e testname        # Single test by name
```

See [docs/QUALITY.md](../../docs/QUALITY.md) for full testing guide.
