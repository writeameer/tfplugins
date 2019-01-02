#!/usr/bin/env bash

# Create output folder
mkdir -p ~/bin

# Build Plugin
go build -o ./bin/plugin ./plugin/main.go
chmod +x ./bin/plugin
# Build Host
go build -o ./bin/main ./main.go
chmod +x ./bin/main

# Run host
cd ./bin
./main
cd ..
