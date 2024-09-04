#!/bin/bash
mkdir -p build/bin
go build $FLAGS -o build/bin ./cmd/...

echo ""
echo "Compiling is complete"

echo ""
echo "Binaries can be found here: ~/anumad/build/bin"