SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)
VER := 0.14

.PHONY: all
all : test build docker

.PHONY: build
build : bin/star-tracker

.PHONY: docker
docker : bin/star-tracker.tar.gz

bin/star-tracker: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" -installsuffix cgo -o ./bin/star-tracker . 
	upx --brute bin/star-tracker

bin/star-tracker.tar.gz: bin/star-tracker star-tracker.json.tmpl
	./build_save.sh

pipeline : star-tracker.json
	pachctl create-pipeline -f star-tracker.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf bin/
	rm -rf test-results
	rm -rf statsapi/
	rm -rf out/
