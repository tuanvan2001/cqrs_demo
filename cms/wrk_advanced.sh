#!/bin/bash

# Advanced Performance test script for CMS API
# Tests different scenarios with varying loads

SERVER_IP="10.73.48.178"
SERVER_PORT="8080"
BASE_URL="http://${SERVER_IP}:${SERVER_PORT}"

echo "=== Advanced CMS API Performance Testing ==="
echo "Server: ${BASE_URL}"
echo "Connection Pool: Max=120, Min=10"
echo ""

# Function to run wrk test with custom parameters
run_custom_test() {
    local name=$1
    local endpoint=$2
    local threads=$3
    local connections=$4
    local duration=$5
    local method=${6:-"GET"}
    local data=${7:-""}
    
    echo "=== ${name} ==="
    echo "Endpoint: ${method} ${endpoint}"
    echo "Load: ${threads} threads, ${connections} connections, ${duration}"
    echo "----------------------------------------"
    
    if [ "$method" = "GET" ]; then
        wrk -t${threads} -c${connections} -d${duration} \
            --latency \
            "${BASE_URL}${endpoint}"
    else
        cat > /tmp/wrk_script.lua << EOF
wrk.method = "${method}"
wrk.body = '${data}'
wrk.headers["Content-Type"] = "application/json"
EOF
        
        wrk -t${threads} -c${connections} -d${duration} \
            --latency \
            -s /tmp/wrk_script.lua \
            "${BASE_URL}${endpoint}"
        
        rm -f /tmp/wrk_script.lua
    fi
    
    echo ""
    echo ""
}

# Check server availability
echo "Checking server availability..."
if ! curl -s --connect-timeout 5 "${BASE_URL}/health" > /dev/null; then
    echo "Server is not accessible at ${BASE_URL}"
    echo "Please start the server first: go run main.go"
    exit 1
fi

echo "Server is running. Starting advanced performance tests..."
echo ""

# Scenario 1: Light Load Test (below connection pool min)
run_custom_test "Light Load - Health Check" "/health" 2 5 "10s"

# Scenario 2: Medium Load Test (around connection pool min)
run_custom_test "Medium Load - Get Locations" "/api/v1/locations" 4 15 "15s"

# Scenario 3: High Load Test (middle range)
run_custom_test "High Load - Get Games" "/api/v1/games" 8 60 "20s"

# Scenario 4: Heavy Load Test (approaching connection pool max)
run_custom_test "Heavy Load - Get News" "/api/v1/news" 12 100 "20s"

# Scenario 5: Maximum Load Test (at connection pool max)
run_custom_test "Maximum Load - Get Settings" "/api/v1/settings" 16 120 "25s"

# Scenario 6: Overload Test (exceeding connection pool max)
run_custom_test "Overload Test - Get Banners" "/api/v1/banners" 20 150 "30s"

# Scenario 7: Write Operations Test
LOCATION_DATA='{"name":"Load Test Location","address":"123 Load Test St","status":1}'
run_custom_test "Write Load - Create Location" "/api/v1/locations" 4 50 "15s" "POST" "$LOCATION_DATA"

# Scenario 8: Mixed Read/Write Test
echo "=== Mixed Read/Write Test ==="
echo "Running parallel read and write operations..."
echo "This will test connection pool efficiency under mixed workload"
echo ""

# Start background read test (heavy load)
echo "Starting background read operations..."
wrk -t8 -c80 -d40s --latency "${BASE_URL}/api/v1/locations" > /tmp/read_results.txt 2>&1 &
READ_PID=$!

# Wait a moment for read test to ramp up
sleep 2

# Start write test (moderate load)
echo "Starting write operations..."
GAME_DATA='{"location_id":1,"name":"Concurrent Test Game","description":"Testing concurrent operations","status":1}'
cat > /tmp/wrk_write.lua << EOF
wrk.method = "POST"
wrk.body = '${GAME_DATA}'
wrk.headers["Content-Type"] = "application/json"
EOF

wrk -t4 -c40 -d35s --latency -s /tmp/wrk_write.lua "${BASE_URL}/api/v1/games" > /tmp/write_results.txt 2>&1 &
WRITE_PID=$!

# Wait for both tests to complete
echo "Waiting for concurrent tests to complete..."
wait $READ_PID
wait $WRITE_PID

echo "Read Test Results:"
echo "==================="
cat /tmp/read_results.txt
echo ""

echo "Write Test Results:"
echo "==================="
cat /tmp/write_results.txt
echo ""

# Cleanup
rm -f /tmp/wrk_write.lua /tmp/read_results.txt /tmp/write_results.txt

# Scenario 9: Sustained Load Test
echo "=== Sustained Load Test ==="
echo "Testing sustained performance over extended period..."
run_custom_test "Sustained Load - Mixed Endpoints" "/api/v1/locations" 8 80 "60s"

# Scenario 10: Burst Test
echo "=== Burst Test ==="
echo "Testing short bursts of high load..."
for i in {1..3}; do
    echo "Burst $i/3:"
    run_custom_test "Burst $i" "/health" 16 200 "5s"
    sleep 2
done

echo ""
echo "=== Advanced Performance Test Complete ==="
echo ""
echo "Connection Pool Benefits Analysis:"
echo "- Light loads (5-15 connections): Should use minimal pool connections efficiently"
echo "- Medium loads (15-60 connections): Should scale within idle connection range"
echo "- Heavy loads (60-120 connections): Should utilize full pool capacity"
echo "- Overload (>120 connections): Should queue requests efficiently without errors"
echo ""
echo "Key Performance Indicators:"
echo "- Consistent latency across different load levels"
echo "- High throughput without connection errors"
echo "- Efficient connection reuse (no connection timeout errors)"
echo "- Graceful handling of loads exceeding pool capacity"
