#!/bin/bash

# 检查GOOS GOARCH 是否设置
if [ -z "$GOOS" ]; then
  export GOOS=$(go env GOOS)
fi

if [ -z "$GOARCH" ]; then
  export GOARCH=$(go env GOARCH)
fi

swag init -d .. -o ../docs

# 打包
mkdir -p ./out
mkdir -p ./temp
rm -rf ./out/*
rm -rf ./temp/*

go build -o ./temp/library_sys_go-${GOOS}-${GOARCH} -trimpath -ldflags="-s -w" ../main.go

cp -r ../docs ./temp/docs

tar -zcvf ./out/library_sys_go-${GOOS}-${GOARCH}.tar.gz -C ./temp/ .
