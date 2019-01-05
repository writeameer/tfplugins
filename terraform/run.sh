#!/bin/bash -e

# Create output folder
mkdir -p ./bin

# Build Plugin
go build -o ./bin/plugin ./plugin/main.go
chmod +x ./bin/plugin
# Build Host
go build -o ./bin/main ./main.go
chmod +x ./bin/main

# Copy azure plugin to bin
cp ./terraform-provider-azurerm ./bin
# Run host
cd ./bin
./main
cd ..
