# Makefile for visions-core (core repo only)

.PHONY: all build test lint clean

all: build

build:
	go build -o visions-core main.go

test:
	go test ./...

lint:
	golangci-lint run || true

clean:
	rm -f visions-core

