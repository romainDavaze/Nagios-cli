SHELL := /bin/bash

build:
	go build
.PHONY: build

install:
	go get ${gobuild_args} ./...
.PHONY: install

test:
	source ./test/init-env.sh && go test -v ./...
.PHONY: test