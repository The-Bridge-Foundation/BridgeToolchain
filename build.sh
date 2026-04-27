#!/bin/sh
go build -tags=llvm19 -o ./build/bridgec ./bridgec/cmd/
go build -o ./build/bridger ./bridger/cmd/
sha256sum build/bridgec
sha256sum build/bridger

echo "Run ./install.sh as root to install"