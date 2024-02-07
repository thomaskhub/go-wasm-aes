#!/usr/bin/bash
# compile the package into wasm and create a dist directory
# which will contain all output files to run wasm in the browser
# This has been tested on ubuntu 22.04 but should run any many other 
# Linux distributions as well
echo "prepare distribution directory"
rm -rf ./dist
mkdir ./dist

echo "compile go package to wasm"
GOARCH=wasm GOOS=js go build -o ./dist/go-wasm-aes.wasm

echo "copy java script wrapper to dist directory"
cp ./go-wasm-aes.js ./dist

#!!! Change this path if go is installed in a different location
cp /usr/local/go/misc/wasm/wasm_exec.js ./dist


echo "packaging done"
