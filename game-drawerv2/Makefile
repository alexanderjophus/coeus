SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)

.PHONY: all
all : test build docker

.PHONY: build
build : bin/game-drawerv2 bin/game-drawerv2-pachd

.PHONY: docker
docker : bin/game-drawerv2-pachd.tar.gz

bin/game-drawerv2: $(SRC)
	go build -o ./bin/game-drawerv2 .

bin/game-drawerv2-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/statsapi/api/v1/game/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/game-drawerv2-pachd .
	upx --brute bin/game-drawerv2-pachd

bin/game-drawerv2-pachd.tar.gz: bin/game-drawerv2-pachd game-drawerv2.jsonnet.tmpl
	./build_save.sh

.PHONY : pipeline
pipeline :
	pachctl delete-pipeline game-drawerv2
	pachctl create-pipeline -f pipelines/game-drawerv2.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf pipelines/
	rm -rf bin/
	rm -rf test
	rm -rf out/
	rm game-drawerv2.jsonnet
