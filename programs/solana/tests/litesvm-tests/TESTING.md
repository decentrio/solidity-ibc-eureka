# Testing Guide

## Setup Complete

The integration test crate has been set up with the following fixes:

### Issues Fixed:
1. **Import errors** - Added `payer_pubkey()` helper method to avoid trait import issues
2. **Dependency versions** - Updated to use workspace dependencies for consistency
3. **Program paths** - Configured correct relative paths for program binaries
4. **Account setup** - Improved account handling in test context

### To Run Tests:

1. **Build the programs first** (if not already built):
   ```bash
   cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana
   cargo build-sbf -p ics26-router
   cargo build-sbf -p ics07-tendermint
   ```

2. **Run the integration tests**:
   ```bash
   cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana
   cargo test -p ibc-solana-integration-tests
   ```

### Test Structure:
- `tests/minimal_test.rs` - Basic compilation check
- `tests/integration_test.rs` - Router and client management tests
- `tests/cross_program_test.rs` - ICS07/ICS26 interaction tests

### Expected Behavior:
- Tests should compile successfully
- Basic tests (like `test_compile_check`) should pass
- Integration tests may fail if programs aren't built or if there are runtime issues
- Use `cargo test -- --nocapture` to see debug output

### Troubleshooting:
1. If compilation fails, check that all dependencies are installed
2. If tests fail to find programs, ensure they're built with `cargo build-sbf`
3. Check the relative paths in `lib.rs` match your directory structure