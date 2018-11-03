SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)

.PHONY: all
all : test build # docker

.PHONY: build
build : bin/highlight-generator bin/highlight-generator-pachd

.PHONY: docker
docker : bin/highlight-generator-pachd.tar.gz

bin/highlight-generator: $(SRC)
	go build -o ./bin/highlight-generator .

bin/highlight-generator-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/statsapi/api/v1/game/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/highlight-generator-pachd .
	# upx --brute bin/highlight-generator-pachd

bin/highlight-generator-pachd.tar.gz: bin/highlight-generator-pachd highlight-generator.jsonnet.tmpl
	./build_save.sh

.PHONY : pipeline
pipeline :
	pachctl delete-pipeline highlight-generator
	pachctl create-pipeline -f pipelines/highlight-generator.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf pipelines/
	rm -rf bin/
	rm -rf test
	rm -rf out/
	rm highlight-generator.jsonnet
