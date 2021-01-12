#!/bin/sh

VERSION=1.15.1.4.mine

# linux_amd64环境
./gfctl build main.go --cgo --CC=x86_64-linux-musl-gcc --CGO_LDFLAGS="-static" --name gfctl --arch amd64 --system linux --version $VERSION

# windows_amd64环境
./gfctl build main.go --cgo --CC=x86_64-w64-mingw32-gcc --name gfctl --arch amd64 --system windows --version $VERSION

# mac_amd64环境
./gfctl build main.go --cgo --name gfctl --arch amd64 --system darwin --version $VERSION

cd ./bin/$VERSION

tar -zcvf gfctl.$VERSION-darwin-amd64.tar.gz darwin_amd64

tar -zcvf gfctl.$VERSION-linux-amd64.tar.gz linux_amd64

#tar -zcvf gfctl.$VERSION-windows-amd64.tar.gz windows_amd64

zip -r gfctl.$VERSION-windows-amd64.zip windows_amd64
