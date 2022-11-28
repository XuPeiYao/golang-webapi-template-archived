#!/bin/bash

rm -f -r ./assets/swagger/docs
go mod download
swag init -g ./cmd/main.go -o ./assets/swagger/docs --ot json,yaml
