#!/bin/bash
# Quick compile check

cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana

echo "Running cargo check on integration tests..."
cargo check -p ibc-solana-integration-tests 2>&1 | tee compile-check.log

if [ ${PIPESTATUS[0]} -eq 0 ]; then
    echo -e "\n✅ Compilation successful!"
    echo "You can now run: cargo test -p ibc-solana-integration-tests"
else
    echo -e "\n❌ Compilation failed. Check compile-check.log for details."
fi