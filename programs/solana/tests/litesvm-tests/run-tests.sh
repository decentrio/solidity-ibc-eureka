#!/bin/bash
# Test runner script

cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana/tests/litesvm-tests

echo "🧪 Running IBC Solana Integration Tests"
echo "======================================="

# First check if it compiles
echo -e "\n📦 Checking compilation..."
cargo check

if [ $? -ne 0 ]; then
    echo "❌ Compilation failed!"
    exit 1
fi

echo -e "\n✅ Compilation successful!"

# Run the tests
echo -e "\n🚀 Running tests..."
cargo test -- --nocapture

echo -e "\n📊 Test Summary:"
echo "If tests passed: The integration test framework is working correctly!"
echo "If tests failed: Check if the Solana programs are built (cargo build-sbf)"