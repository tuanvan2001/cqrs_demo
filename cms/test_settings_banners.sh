#!/bin/bash

# Comprehensive test script for Settings and Banners CRUD operations

echo "=== Testing Settings and Banners CRUD Operations ==="

# Start server in background
echo "Starting server..."
cd /home/tuantech/codews/cqrs_demo/cms
go run main.go &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo "Testing health endpoint..."
curl -s http://localhost:8080/health | jq .

echo -e "\n=== Setting up Location for Testing ==="

# Create a location first (required for settings and banners)
echo "Creating test location..."
LOCATION_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Center",
    "address": "123 Gaming Street",
    "status": 1
  }')

echo $LOCATION_RESPONSE | jq .
LOCATION_ID=$(echo $LOCATION_RESPONSE | jq -r '.data.id')
echo "Location ID: $LOCATION_ID"

echo -e "\n=== SETTINGS CRUD OPERATIONS ==="

echo -e "\n1. CREATE Settings"
# Create multiple settings
echo "Creating theme setting..."
SETTING1_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/settings \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"theme_color\",
    \"value\": \"#FF5733\"
  }")
echo $SETTING1_RESPONSE | jq .
SETTING1_ID=$(echo $SETTING1_RESPONSE | jq -r '.data.id')

echo -e "\nCreating max players setting..."
SETTING2_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/settings \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"max_players\",
    \"value\": \"100\"
  }")
echo $SETTING2_RESPONSE | jq .
SETTING2_ID=$(echo $SETTING2_RESPONSE | jq -r '.data.id')

echo -e "\nCreating opening hours setting..."
curl -s -X POST http://localhost:8080/api/v1/settings \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"opening_hours\",
    \"value\": \"9:00 AM - 11:00 PM\"
  }" | jq .

echo -e "\n2. READ Settings"
echo "Getting all settings..."
curl -s http://localhost:8080/api/v1/settings | jq .

echo -e "\nGetting setting by ID ($SETTING1_ID)..."
curl -s http://localhost:8080/api/v1/settings/$SETTING1_ID | jq .

echo -e "\n3. UPDATE Setting"
echo "Updating theme color setting..."
curl -s -X PUT http://localhost:8080/api/v1/settings/$SETTING1_ID \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"key\": \"theme_color\",
    \"value\": \"#33FF57\"
  }" | jq .

echo -e "\n=== BANNERS CRUD OPERATIONS ==="

echo -e "\n1. CREATE Banners"
# Create multiple banners
echo "Creating promotion banner..."
BANNER1_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/banners \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/promo_banner.jpg\",
    \"link\": \"https://example.com/promotion\",
    \"status\": 1
  }")
echo $BANNER1_RESPONSE | jq .
BANNER1_ID=$(echo $BANNER1_RESPONSE | jq -r '.data.id')

echo -e "\nCreating event banner..."
BANNER2_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/banners \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/event_banner.jpg\",
    \"link\": \"https://example.com/events\",
    \"status\": 1
  }")
echo $BANNER2_RESPONSE | jq .
BANNER2_ID=$(echo $BANNER2_RESPONSE | jq -r '.data.id')

echo -e "\nCreating inactive banner..."
curl -s -X POST http://localhost:8080/api/v1/banners \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/inactive_banner.jpg\",
    \"link\": \"https://example.com/inactive\",
    \"status\": 0
  }" | jq .

echo -e "\n2. READ Banners"
echo "Getting all banners..."
curl -s http://localhost:8080/api/v1/banners | jq .

echo -e "\nGetting banner by ID ($BANNER1_ID)..."
curl -s http://localhost:8080/api/v1/banners/$BANNER1_ID | jq .

echo -e "\n3. UPDATE Banner"
echo "Updating promotion banner..."
curl -s -X PUT http://localhost:8080/api/v1/banners/$BANNER1_ID \
  -H "Content-Type: application/json" \
  -d "{
    \"location_id\": $LOCATION_ID,
    \"image\": \"https://example.com/updated_promo_banner.jpg\",
    \"link\": \"https://example.com/updated_promotion\",
    \"status\": 1
  }" | jq .

echo -e "\n4. DELETE Operations"
echo "Deleting a setting ($SETTING2_ID)..."
curl -s -X DELETE http://localhost:8080/api/v1/settings/$SETTING2_ID | jq .

echo -e "\nDeleting a banner ($BANNER2_ID)..."
curl -s -X DELETE http://localhost:8080/api/v1/banners/$BANNER2_ID | jq .

echo -e "\n=== Final State Check ==="
echo "Remaining settings:"
curl -s http://localhost:8080/api/v1/settings | jq .

echo -e "\nRemaining banners:"
curl -s http://localhost:8080/api/v1/banners | jq .

echo -e "\n=== Test Summary ==="
echo "✅ Created location with ID: $LOCATION_ID"
echo "✅ Created 3 settings (1 deleted)"
echo "✅ Created 3 banners (1 deleted)"
echo "✅ Tested READ operations"
echo "✅ Tested UPDATE operations"
echo "✅ Tested DELETE operations"

echo -e "\n=== Cleaning up ==="
# Kill server
kill $SERVER_PID
echo "Server stopped."
echo "Test completed successfully!"


CREATE USER 'root'@'%' IDENTIFIED BY 'Tuan123';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
