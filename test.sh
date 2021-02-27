#!/bin/bash
go build
for filename in ./testdata/*.txt; do
    echo "$filename:"
    ./expertsystem "$filename"
    echo
done
