#!/bin/bash

echo "Checking LiteSVM test crate structure..."

# Check if Cargo.toml exists
if [ -f "Cargo.toml" ]; then
    echo "✓ Cargo.toml found"
else
    echo "✗ Cargo.toml missing"
fi

# Check source files
for file in "src/lib.rs" "src/fixtures.rs" "src/utils.rs"; do
    if [ -f "$file" ]; then
        echo "✓ $file found"
    else
        echo "✗ $file missing"
    fi
done

# Check test files
for file in "tests/integration_test.rs" "tests/cross_program_test.rs"; do
    if [ -f "$file" ]; then
        echo "✓ $file found"
    else
        echo "✗ $file missing"
    fi
done

echo ""
echo "Dependencies check:"
grep -E "ics07-tendermint|ics26-router|litesvm" Cargo.toml | head -10

echo ""
echo "Workspace inclusion check:"
grep -q "tests/litesvm-tests" ../../Cargo.toml && echo "✓ Included in workspace" || echo "✗ Not in workspace"