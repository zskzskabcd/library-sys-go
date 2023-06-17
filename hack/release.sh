#!/bin/bash

OS_LIST="darwin linux windows"
ARCH_LIST="amd64 arm64 386"

for os in $OS_LIST; do
  for arch in $ARCH_LIST; do
    echo "Building for $os-$arch"
    GOOS=$os GOARCH=$arch bash ./releaser.sh
  done
done