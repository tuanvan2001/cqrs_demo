#!/bin/bash

# Test script for Go CRUD API

echo "=== Testing Go CRUD API ==="

# Start server in background
echo "Starting server..."
go run main.go &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo "Testing health endpoint..."
curl -s http://localhost:8080/health | jq .

echo -e "\n=== Testing Locations API ==="

# Create location
echo "Creating location..."
LOCATION_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Location",
    "address": "123 Test Street",
    "status": 1
  }')

echo $LOCATION_RESPONSE | jq .

LOCATION_ID=$(echo $LOCATION_RESPONSE | jq -r '.data.id')

# Get all locations
echo -e "\nGetting all locations..."
curl -s http://localhost:8080/api/v1/locations | jq .

echo -e "\n=== Testing Games API ==="

# Create game
echo "Creating game..."
GAME_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/games \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"name\": \"Test Game\",
    \"description\": \"This is a test game\",
    \"status\": 1
  }")

echo $GAME_RESPONSE | jq .

# Get all games
echo -e "\nGetting all games..."
curl -s http://localhost:8080/api/v1/games | jq .

echo -e "\n=== Testing News API ==="

# Create news
echo "Creating news..."
NEWS_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/news \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test News",
    "content": "This is test news content",
    "status": 1
  }')

echo $NEWS_RESPONSE | jq .

# Get all news
echo -e "\nGetting all news..."
curl -s http://localhost:8080/api/v1/news | jq .

echo -e "\n=== Testing Settings API ==="

# Create setting
echo "Creating setting..."
SETTING_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/settings \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"theme_color\",
    \"value\": \"#FF5733\"
  }")

echo $SETTING_RESPONSE | jq .

SETTING_ID=$(echo $SETTING_RESPONSE | jq -r '.data.id')

# Create another setting
echo -e "\nCreating another setting..."
curl -s -X POST http://localhost:8080/api/v1/settings \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"max_players\",
    \"value\": \"100\"
  }" | jq .

# Get all settings
echo -e "\nGetting all settings..."
curl -s http://localhost:8080/api/v1/settings | jq .

# Get setting by ID
echo -e "\nGetting setting by ID..."
curl -s http://localhost:8080/api/v1/settings/$SETTING_ID | jq .

echo -e "\n=== Testing Banners API ==="

# Create banner
echo "Creating banner..."
BANNER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/banners \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/banner1.jpg\",
    \"link\": \"https://example.com/promotion1\",
    \"status\": 1
  }")

echo $BANNER_RESPONSE | jq .

BANNER_ID=$(echo $BANNER_RESPONSE | jq -r '.data.id')

# Create another banner
echo -e "\nCreating another banner..."
curl -s -X POST http://localhost:8080/api/v1/banners \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/banner2.jpg\",
    \"link\": \"https://example.com/promotion2\",
    \"status\": 1
  }" | jq .

# Get all banners
echo -e "\nGetting all banners..."
curl -s http://localhost:8080/api/v1/banners | jq .

# Get banner by ID
echo -e "\nGetting banner by ID..."
curl -s http://localhost:8080/api/v1/banners/$BANNER_ID | jq .

echo -e "\n=== Testing Update Operations ==="

# Update location
echo "Updating location..."
curl -s -X PUT http://localhost:8080/api/v1/locations/$LOCATION_ID \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Test Location",
    "address": "456 Updated Street",
    "status": 1
  }' | jq .

# Update setting
echo -e "\nUpdating setting..."
curl -s -X PUT http://localhost:8080/api/v1/settings/$SETTING_ID \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"theme_color\",
    \"value\": \"#33FF57\"
  }" | jq .

# Update banner
echo -e "\nUpdating banner..."
curl -s -X PUT http://localhost:8080/api/v1/banners/$BANNER_ID \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/updated_banner.jpg\",
    \"link\": \"https://example.com/updated_promotion\",
    \"status\": 0
  }" | jq .

echo -e "\n=== Summary ==="
echo "Created records:"
echo "- Location ID: $LOCATION_ID"
echo "- Game ID: $(echo $GAME_RESPONSE | jq -r '.data.id')"
echo "- News ID: $(echo $NEWS_RESPONSE | jq -r '.data.id')"
echo "- Setting ID: $SETTING_ID"
echo "- Banner ID: $BANNER_ID"

echo -e "\n=== Cleaning up ==="
# Kill server
kill $SERVER_PID
echo "Server stopped."
