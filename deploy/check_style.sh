#!/bin/bash

# Check usage.
if [ $# -ne 1 ]; then
    echo "USAGE: check_style.sh <source directory>"
    exit 1
fi

# Find any files with goimports problems
IMP_FILES=$(goimports -e -l -w $(find $1 -wholename "*.go" -not -wholename "*/Godeps/*"))

if [ -n "$IMP_FILES" ]; then
    echo "The following files are not properly formatted:"
    echo ${IMP_FILES}
    exit 1
fi

echo "All file are properly formatted."
