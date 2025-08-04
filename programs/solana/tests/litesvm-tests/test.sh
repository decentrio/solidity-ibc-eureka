#!/bin/bash
# Simple test script to verify the integration tests can compile

cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana

# Check if the test crate compiles
echo "Checking if test crate compiles..."
cargo check -p ibc-solana-integration-tests

# Run the tests
echo "Running integration tests..."
cargo test -p ibc-solana-integration-tests

echo "Test run complete!"