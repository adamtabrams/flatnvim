#!/bin/sh

go build -o bin/ || { echo "failure"; exit 1; }
echo "success"
echo "path to the flatnvim binary: $PWD/bin/flatnvim"
