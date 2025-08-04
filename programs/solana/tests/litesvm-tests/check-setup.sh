#!/bin/bash
# Script to check test setup

echo "Checking test crate setup..."
echo "============================"

echo -e "\n1. Checking if we're in the right directory:"
pwd

echo -e "\n2. Checking Cargo.toml exists:"
ls -la Cargo.toml

echo -e "\n3. Checking source files:"
find src -name "*.rs" | head -10

echo -e "\n4. Checking test files:"
find tests -name "*.rs" | head -10

echo -e "\n5. Checking if target directory exists:"
ls -la ../../../../target/deploy/ 2>/dev/null || echo "Target directory not found - programs need to be built first"

echo -e "\n6. Workspace Cargo.toml location:"
ls -la ../../Cargo.toml

echo -e "\nTo run tests:"
echo "cd /Users/vaporif/Repos/solidity-ibc-eureka/programs/solana"
echo "cargo test -p ibc-solana-integration-tests"

echo -e "\nTo build programs first:"
echo "cargo build-sbf -p ics26-router"
echo "cargo build-sbf -p ics07-tendermint"