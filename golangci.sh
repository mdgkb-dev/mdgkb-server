#!/bin/bash
echo Test1
GOLANGCI="golangci-lint"
opt="run -E gofmt -E goimports -E golint  -E asciicheck -E unparam -E gosec -E typecheck -E unconvert"

$GOLANGCI $opt ./...