#!/bin/bash

set -e

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o out/main cmd/api/main.go" --command="./out/main"