#!/bin/sh

ls
cd /go/src/goChess

go clean
go install ./...
ls

go build
./goChess
