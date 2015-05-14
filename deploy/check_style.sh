#!/bin/bash

# Check usage.
if [ $# -ne 1 ]; then
  echo "USAGE: check_goimports <source directory>"
  exit 1
fi

# Check formatting on non Godep'd code.
GOFMT_PATHS=$(find . -not -wholename "*.git*" -not -wholename "*Godeps*" -not -name "." -type d)

# Find any files with gofgomt problems
FMT_FILES=$(gofmt -s -l ${GOFMT_PATHS})

if [ -n "$FMT_FILES" ]; then
  echo "The following files are not properly formatted:"
  echo ${FMT_FILES}
  exit 1
fi

# Find any files with goimports problems
IMP_FILES=$(goimports -e -l ${GOFMT_PATHS})

if [ -n "$IMP_FILES" ]; then
  echo "The following files have imports which are not properly formatted:"
  echo ${IMP_FILES}
  exit 1
fi