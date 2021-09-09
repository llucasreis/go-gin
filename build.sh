#!/usr/bin/bash
set -xe

# install packages
go get github.com/gin-gonic/gin

go get github.com/go-playground/validator/v10

# build command
go build -o bin/application server.go