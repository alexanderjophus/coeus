SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)

.PHONY: all
all : test build docker

.PHONY: build
build : bin/game-drawer bin/game-drawer-pachd

.PHONY: docker
docker : bin/game-drawer-pachd.tar.gz

bin/game-drawer: $(SRC)
	go build -o ./bin/game-drawer .

bin/game-drawer-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/play-tracker/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/game-drawer-pachd .
	upx --brute bin/game-drawer-pachd

bin/game-drawer-pachd.tar.gz: bin/game-drawer-pachd game-drawer.json.tmpl
	./build_save.sh

pipeline : game-drawer.json
	pachctl delete-pipeline game-drawer
	pachctl create-pipeline -f game-drawer.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf bin/
	rm -rf test
	rm -rf out/
	rm game-drawer.json
