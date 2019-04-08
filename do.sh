#!/bin/sh

APP_NAME=${APP_NAME:-echo-boiler}
ENV=${ENV:-local}

start () {
    export ENV=$ENV
    export APP_NAME=$APP_NAME
    go build -v -o ./bin/$APP_NAME ./cmd/service/main.go && ./bin/$APP_NAME
}

run_tests () {
    go test -v ./... -cover
}

run_bench () {
    go test -v ./... -bench Benchmark
}

$*