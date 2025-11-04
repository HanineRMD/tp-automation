#!/bin/bash

echo "🧪 LANCEMENT DES TESTS COMPLETS"
echo "================================"

# Test Python
echo ""
echo "🐍 TEST SCRIPTS PYTHON"
for script in python/ex*.py; do
    echo "Testing $script..."
    python3 "$script" --help 2>/dev/null && echo "✅ $script OK" || echo "❌ $script FAILED"
done

# Test Shell
echo ""
echo "🐚 TEST SCRIPTS SHELL"
for script in shell/ex*.sh; do
    echo "Testing $script..."
    bash -n "$script" && echo "✅ $script OK" || echo "❌ $script FAILED"
done

# Test Go
echo ""
echo "🦫 TEST PROGRAMMES GO"
for script in go/ex*.go; do
    echo "Testing $script..."
    go build -o /tmp/test_binary "$script" 2>/dev/null && echo "✅ $script OK" || echo "❌ $script FAILED"
done

echo ""
echo "🎯 TOUS LES TESTS TERMINÉS!"