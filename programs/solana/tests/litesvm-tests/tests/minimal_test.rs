//! Minimal test to verify setup

#[test]
fn test_basic_setup() {
    // Just verify the test crate compiles and can run a basic test
    assert_eq!(2 + 2, 4);
}

#[test]
fn test_imports() {
    // Verify we can import from the test crate
    use ibc_solana_litesvm_tests::IbcTestContext;
    
    // This will only compile if all dependencies are correctly set up
    let _ctx = IbcTestContext::new();
}