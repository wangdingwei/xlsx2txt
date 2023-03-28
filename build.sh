#!/usr/bin/env bash

cd "$(dirname $0)"

mkdir -p build

build() {
    local os=$(go env GOOS)
    local arch=$(go env GOARCH)
    local suffix=""
    if [[ $os == win* ]]; then
        suffix=.exe
    fi
    go build -o build/xlsx2txt_${os}_${arch}${suffix}
}

GOOS=windows GOARCH=amd64 build
GOOS=darwin GOARCH=amd64 build
GOOS=darwin GOARCH=arm64 build
