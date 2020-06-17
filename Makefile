SHELL := /bin/bash

BASEDIR = $(shell pwd)

CURRENTNO := 10

all: run

run:
	@go run ${CURRENTNO}\#.go
test:
	@go test -v $(CURRENTNO)\#.go $(CURRENTNO)\#_test.go