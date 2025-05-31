#!/bin/bash

# Script tạo nhiều location cho CMS Redis qua API

API_URL="http://localhost:8081/api/v1/locations"

echo "Tạo location 1..."
curl -s -X POST "$API_URL" -H "Content-Type: application/json" -d '{
  "name": "Location 1",
  "address": "123 Main St",
  "status": 1
}'

echo -e "\nTạo location 2..."
curl -s -X POST "$API_URL" -H "Content-Type: application/json" -d '{
  "name": "Location 2",
  "address": "456 Second Ave",
  "status": 1
}'

echo -e "\nTạo location 3..."
curl -s -X POST "$API_URL" -H "Content-Type: application/json" -d '{
  "name": "Location 3",
  "address": "789 Third Blvd",
  "status": 0
}'

echo -e "\nHoàn tất tạo location mẫu!"