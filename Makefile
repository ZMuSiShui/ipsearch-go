SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP_NAME=ipsearch
APP_VERSION=1.1
IMAGE_NAME="ZMuSiShui/${APP_NAME}:${APP_VERSION}"
IMAGE_LATEST="ZMuSiShui/${APP_NAME}:latest"

all: mod fmt imports lint test
first:
	mkdir src
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
fmt:
	gofmt -w .
mod:
	go mod tidy
imports:
	goimports -w .
lint:
	golangci-lint run
.PHONY: build
build:
	rm ipsearch
	go build -o ipsearch cmd/app.go
build-mac:
	rm -f ipsearch ipsearch-darwin-amd64.tar.gz
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ipsearch cmd/app.go
	tar -zcvf src/ipsearch-darwin-amd64.tar.gz ipsearch
build-linux:
	rm -f ipsearch ipsearch-linux-amd64.tar.gz
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ipsearch cmd/app.go
	tar -zcvf src/ipsearch-linux-amd64.tar.gz ipsearch
build-win:
	rm -f ipsearch.exe ipsearch-windows-amd64.tar.gz
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ipsearch.exe cmd/app.go
	tar -zcvf src/ipsearch-windows-amd64.tar.gz ipsearch.exe
build-win32:
	rm -f ipsearch.exe ipsearch-windows-386.tar.gz
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ipsearch.exe cmd/app.go
	tar -zcvf src/ipsearch-windows-386.tar.gz ipsearch.exe
build-release:
	make build-mac
	make build-linux
	make build-win
	make build-win32
	rm -f ipsearch ipsearch.exe
help:
	@echo "first - first time"
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "imports - go imports"
	@echo "lint - run golangci-lint"
	@echo "build - build binary"
	@echo "build-mac - build mac binary"
	@echo "build-linux - build linux amd64 binary"
	@echo "build-win - build win amd64 binary"
	@echo "build-win32 - build win 386 binary"