#!/bin/sh

which go >/dev/null ||
    { echo "cannot build without 'go' installed"; exit 2; }

go build -o bin/ || { echo "failure"; exit 1; }
echo "build successful"
echo "path to the flatnvim binary: $PWD/bin/flatnvim"
