.PHONY: all test clean build install

all: install test

default:
    go build ./...

build:
	go build ./...

install:
	go get ./...

test: install
	go test ./...

clean:
	go clean -i ./...