//! Simple test to check if the crate compiles

use ibc_solana_litesvm_tests::IbcTestContext;

#[test]
fn test_compile_check() {
    // Just check that we can create context and access pubkey
    let ctx = IbcTestContext::new();
    let _pubkey = ctx.payer_pubkey();
    
    // If this compiles and runs, our imports are correct
    assert_eq!(1 + 1, 2);
}