#!/bin/sh

go build -o bin/ || { echo "failure"; exit 1; }
echo "build successful"
echo "path to the flatnvim binary: $PWD/bin/flatnvim"
