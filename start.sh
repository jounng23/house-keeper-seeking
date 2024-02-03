#!/bin/bash

# Run DB container
docker-compose up -d &

# Run service 1
cd pricing-service/
go mod tidy &
go run main.go &

# Run service 2
cd ../ 
cd sending-service/
go mod tidy &
go run main.go &

# Run service 3
cd ../
cd booking-service/
go mod tidy &
go run main.go &