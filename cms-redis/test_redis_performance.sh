#!/bin/bash

# Redis Performance Test Script
# Tests Redis connection pool optimization and read performance

echo "=== Redis Performance Test ==="
echo "Testing optimized Redis connection pool configuration"
echo ""

SERVER_IP="10.73.48.178"
SERVER_PORT="8081"
BASE_URL="http://${SERVER_IP}:${SERVER_PORT}"

# Check if wrk is installed
if ! command -v wrk &> /dev/null; then
    echo "Installing wrk for performance testing..."
    sudo apt-get update && sudo apt-get install -y wrk
fi

# Function to run performance test
run_redis_test() {
    local name=$1
    local endpoint=$2
    local threads=$3
    local connections=$4
    local duration=$5
    
    echo "=== ${name} ==="
    echo "Endpoint: GET ${endpoint}"
    echo "Load: ${threads} threads, ${connections} connections, ${duration}"
    echo "Testing Redis connection pool performance..."
    echo "----------------------------------------"
    
    wrk -t${threads} -c${connections} -d${duration} \
        --latency \
        "${BASE_URL}${endpoint}"
    
    echo ""
    echo ""
}

# Check server availability
echo "Checking Redis server availability..."
if ! curl -s --connect-timeout 5 "${BASE_URL}/api/v1/locations" > /dev/null; then
    echo "Server is not accessible at ${BASE_URL}"
    echo "Please start the Redis server first: cd cms-redis && go run main.go"
    exit 1
fi

echo "Redis server is running. Starting performance tests..."
echo ""

# Test 1: Light load - below pool minimum
run_redis_test "Light Load Redis Read" "/api/v1/locations" 5 15 "30s"

# Test 2: Medium load - within idle connection range
run_redis_test "Medium Load Redis Read" "/api/v1/locations" 10 30 "30s"

# Test 3: Heavy load - using full pool capacity
run_redis_test "Heavy Load Redis Read" "/api/v1/locations" 20 60 "30s"

# Test 4: Overload - testing pool queue handling
run_redis_test "Overload Redis Read" "/api/v1/locations" 30 120 "30s"

# Test 5: Single key read performance
run_redis_test "Single Key Read Performance" "/api/v1/locations/1" 15 50 "30s"

echo "=== Redis Performance Test Complete ==="
echo ""
echo "Redis Connection Pool Benefits Analysis:"
echo "✓ Pool Size: 100 - Maximum concurrent connections"
echo "✓ Min Idle: 20 - Always ready connections for fast response"
echo "✓ Max Idle: 50 - Optimal connection reuse"
echo "✓ Connection Lifetime: 1h - Prevents connection leaks"
echo "✓ Read Timeout: 3s - Fast timeout for read operations"
echo "✓ Pool Timeout: 4s - Quick connection acquisition"
echo "✓ Retry Logic: 3 retries with exponential backoff"
echo ""
echo "Expected Performance Improvements:"
echo "• Reduced Redis connection establishment overhead"
echo "• Better connection reuse and pooling"
echo "• Faster read operations with optimized timeouts"
echo "• Improved handling of concurrent Redis requests"
echo "• Graceful retry handling for temporary failures"
echo ""
echo "Monitor Redis connection usage:"
echo "redis-cli INFO clients"
echo "redis-cli CLIENT LIST"
