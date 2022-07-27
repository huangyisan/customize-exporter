.PHONY: file-linux file-osx file-windows clean

PWD=$(shell pwd)

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

file-linux:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=linux go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.mode=file" -o bin/files-exporter

file-osx:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.mode=file" -o bin/files-exporter

file-windows:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build -ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.mode=file" -o bin/files-exporter

clean:
	@rm -rf bin/*