#! /bin/bash
export GOPATH="$(pwd)"
go get github.com/atotto/clipboard
go build kv.go
