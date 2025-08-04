# IBC Solana Integration Tests

This crate provides integration tests for the IBC Solana programs using `mollusk-svm`.

## Structure

- `src/lib.rs` - Main test context and utilities
- `src/fixtures.rs` - Test data and fixtures
- `src/utils.rs` - Helper functions for creating instructions and PDAs
- `tests/integration_test.rs` - Basic integration tests
- `tests/cross_program_test.rs` - Cross-program interaction tests

## Running Tests

```bash
cargo test
```

## Test Context

The `IbcTestContext` provides a test environment with:
- Mollusk SVM for simulating Solana runtime
- Pre-configured program IDs for ICS07 and ICS26
- Helper methods for processing instructions

## Example Usage

```rust
use ibc_solana_litesvm_tests::{fixtures::*, utils::*, IbcTestContext};

#[test]
fn test_router_initialization() {
    let mut ctx = IbcTestContext::new();
    
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer.pubkey(),
        &ctx.payer.pubkey(),
    );
    
    let result = ctx.process_transaction(vec![init_ix]);
    assert!(result.program_result.is_ok());
}
```