#!/bin/bash

set -e

GOOS=linux GOARCH=amd64 go build -o skeeter-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o skeeter-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o skeeter-windows-amd64.exe
