#!/bin/sh

# 安装交叉编译工具
# brew install mingw-w64
# brew install FiloSottile/musl-cross/musl-cross

# 编译linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -v -o gfctl-1.15.0.2.mine-linux-amd64

# 编译windows
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -v -o gfctl-1.15.0.2.mine-windows-amd64.exe

# 编译mac
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o gfctl-1.15.0.2.mine-darwin-amd64
