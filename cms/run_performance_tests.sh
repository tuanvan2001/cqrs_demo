#!/bin/bash

# Main script to run all performance tests with connection pool optimization

echo "=== CMS API Performance Test Suite ==="
echo "MySQL Connection Pool: Max=120, Min=10, Lifetime=1h"
echo ""

# Check if wrk is installed
if ! command -v wrk &> /dev/null; then
    echo "Installing wrk..."
    
    # Detect OS and install wrk
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get &> /dev/null; then
            sudo apt-get update && sudo apt-get install -y wrk
        elif command -v yum &> /dev/null; then
            sudo yum install -y wrk
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y wrk
        else
            echo "Please install wrk manually:"
            echo "Visit: https://github.com/wg/wrk"
            exit 1
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        if command -v brew &> /dev/null; then
            brew install wrk
        else
            echo "Please install Homebrew first, then run: brew install wrk"
            exit 1
        fi
    fi
fi

# Make scripts executable
chmod +x wrk.sh
chmod +x wrk_advanced.sh

echo "Available test types:"
echo "1. Basic Performance Test (wrk.sh) - Standard load testing"
echo "2. Advanced Performance Test (wrk_advanced.sh) - Connection pool stress testing"
echo "3. Both tests - Complete performance analysis"
echo "4. Quick test - Fast validation"

read -p "Enter choice (1/2/3/4): " choice

case $choice in
    1)
        echo "Running basic performance test..."
        echo "This will test all API endpoints with standard load (4 threads, 500 connections, 20s)"
        echo ""
        ./wrk.sh
        ;;
    2)
        echo "Running advanced performance test..."
        echo "This will test connection pool efficiency with varying loads"
        echo ""
        ./wrk_advanced.sh
        ;;
    3)
        echo "Running complete performance analysis..."
        echo ""
        echo "=== Phase 1: Basic Performance Test ==="
        ./wrk.sh
        echo ""
        echo "=== Phase 2: Advanced Connection Pool Test ==="
        ./wrk_advanced.sh
        ;;
    4)
        echo "Running quick validation test..."
        SERVER_IP="10.73.48.178"
        SERVER_PORT="8080"
        BASE_URL="http://${SERVER_IP}:${SERVER_PORT}"
        
        # Quick health check
        echo "Testing server connectivity..."
        wrk -t2 -c10 -d5s --latency "${BASE_URL}/health"
        
        echo ""
        echo "Quick API test..."
        wrk -t4 -c50 -d10s --latency "${BASE_URL}/api/v1/locations"
        ;;
    *)
        echo "Invalid choice"
        exit 1
        ;;
esac

echo ""
echo "=== Performance Testing Summary ==="
echo ""
echo "Connection Pool Configuration Benefits:"
echo "✓ Max Open Connections: 120 - Handles high concurrent load"
echo "✓ Max Idle Connections: 10 - Maintains ready connections for quick response"
echo "✓ Connection Lifetime: 1 hour - Prevents connection leaks and timeouts"
echo ""
echo "Expected Performance Improvements:"
echo "• Reduced connection establishment overhead"
echo "• Better resource utilization under load"
echo "• Improved response times for concurrent requests"
echo "• Graceful handling of traffic spikes"
echo ""
echo "Performance testing completed!"
echo "Monitor MySQL connection usage with: SHOW PROCESSLIST; or SHOW STATUS LIKE 'Threads_connected';"
