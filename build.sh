#!/bin/sh

# 编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gfctl-1.15.0-linux-amd64

# 编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gfctl-1.15.0-windows-amd64.exe

# 编译mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gfctl-1.15.0-darwin-amd64
