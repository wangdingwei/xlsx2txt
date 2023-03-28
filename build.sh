#!/usr/bin/env bash

cd "$(dirname $0)"

mkdir -p build

build() {
    local os=$(go env GOOS)
    local arch=$(go env GOARCH)
    local file=xlsx2txt
    if [[ $os == win* ]]; then
        file=$file.exe
    fi
    go build -o build/$file
    (
        cd build
        zip xlsx2txt_${os}_${arch}.zip $file
        rm -f $file
    )
    
    
}

GOOS=windows GOARCH=amd64 build
GOOS=darwin GOARCH=amd64 build
GOOS=darwin GOARCH=arm64 build
