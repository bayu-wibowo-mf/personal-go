#!/bin/bash

echo "Starting API Stress Tests..."

# Test 1: GET /employeeDemography - List performance
echo "\nTesting GET /employeeDemography endpoint (100 concurrent users, 1000 requests)"
ab -n 1000 -c 100 http://localhost:8080/employeeDemography/
