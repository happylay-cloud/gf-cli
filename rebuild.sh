#!/bin/sh

# linux_amd64环境
./gfctl build main.go --cgo --CC=x86_64-linux-musl-gcc --CGO_LDFLAGS="-static" --name gfctl --arch amd64 --system linux --version 1.15.0.3.mine

# windows_amd64环境
./gfctl build main.go --cgo --CC=x86_64-w64-mingw32-gcc --name gfctl --arch amd64 --system windows --version 1.15.0.3.mine

# mac_amd64环境
./gfctl build main.go --cgo --name gfctl --arch amd64 --system darwin --version 1.15.0.3.mine

