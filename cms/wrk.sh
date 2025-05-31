#!/bin/bash

# Performance test script using wrk for CMS API
# Usage: ./wrk.sh
# Requirements: wrk tool must be installed

SERVER_IP="10.73.48.178"
SERVER_PORT="8080"
BASE_URL="http://${SERVER_IP}:${SERVER_PORT}"

# Test parameters
THREADS=4
CONNECTIONS=500
DURATION="20s"

echo "=== CMS API Performance Testing ==="
echo "Server: ${BASE_URL}"
echo "Threads: ${THREADS}, Connections: ${CONNECTIONS}, Duration: ${DURATION}"
echo ""

# Function to run wrk test
run_wrk_test() {
    local endpoint=$1
    local method=${2:-"GET"}
    local data=${3:-""}
    local content_type=${4:-"application/json"}
    
    echo "Testing: ${method} ${endpoint}"
    echo "----------------------------------------"
    
    if [ "$method" = "GET" ]; then
        wrk -t${THREADS} -c${CONNECTIONS} -d${DURATION} \
            --latency \
            "${BASE_URL}${endpoint}"
    else
        # Create Lua script for POST/PUT requests
        cat > /tmp/wrk_script.lua << EOF
wrk.method = "${method}"
wrk.body = '${data}'
wrk.headers["Content-Type"] = "${content_type}"
EOF
        
        wrk -t${THREADS} -c${CONNECTIONS} -d${DURATION} \
            --latency \
            -s /tmp/wrk_script.lua \
            "${BASE_URL}${endpoint}"
        
        rm -f /tmp/wrk_script.lua
    fi
    
    echo ""
    echo ""
}

# Check if wrk is installed
if ! command -v wrk &> /dev/null; then
    echo "wrk is not installed. Please install it first:"
    echo "Ubuntu/Debian: sudo apt-get install wrk"
    echo "CentOS/RHEL: sudo yum install wrk"
    echo "macOS: brew install wrk"
    exit 1
fi

# Check if server is running
echo "Checking server availability..."
if ! curl -s --connect-timeout 5 "${BASE_URL}/health" > /dev/null; then
    echo "Server is not running or not accessible at ${BASE_URL}"
    echo "Please start the server first: go run main.go"
    exit 1
fi

echo "Server is running. Starting performance tests..."
echo ""

# Test 1: Health Check
run_wrk_test "/health"

# Test 2: Get all locations
run_wrk_test "/api/v1/locations"

# Test 3: Get all games
run_wrk_test "/api/v1/games"

# Test 4: Get all news
run_wrk_test "/api/v1/news"

# Test 5: Get all settings
run_wrk_test "/api/v1/settings"

# Test 6: Get all banners
run_wrk_test "/api/v1/banners"

# Test 7: Create location (POST)
LOCATION_DATA='{
    "name": "Performance Test Location",
    "address": "123 Performance Street",
    "status": 1
}'
run_wrk_test "/api/v1/locations" "POST" "$LOCATION_DATA"

# Test 8: Create game (POST)
GAME_DATA='{
    "location_id": 1,
    "name": "Performance Test Game",
    "description": "This is a performance test game",
    "status": 1
}'
run_wrk_test "/api/v1/games" "POST" "$GAME_DATA"

# Test 9: Create news (POST)
NEWS_DATA='{
    "title": "Performance Test News",
    "content": "This is performance test news content",
    "status": 1
}'
run_wrk_test "/api/v1/news" "POST" "$NEWS_DATA"

# Test 10: Create setting (POST)
SETTING_DATA='{
    "location_id": 1,
    "key": "test_setting",
    "value": "test_value"
}'
run_wrk_test "/api/v1/settings" "POST" "$SETTING_DATA"

# Test 11: Create banner (POST)
BANNER_DATA='{
    "location_id": 1,
    "image": "https://example.com/test_banner.jpg",
    "link": "https://example.com/test_link",
    "status": 1
}'
run_wrk_test "/api/v1/banners" "POST" "$BANNER_DATA"

echo "=== Performance Test Summary ==="
echo "All tests completed!"
echo "Check the results above for detailed performance metrics."
echo ""
echo "Key metrics to analyze:"
echo "- Requests/sec: Higher is better"
echo "- Latency: Lower is better"
echo "- Transfer/sec: Data throughput"
echo "- Error rate: Should be 0% or very low"
echo ""
echo "Connection Pool Configuration:"
echo "- Max Open Connections: 120"
echo "- Max Idle Connections: 10"
echo "- Connection Max Lifetime: 1 hour"
