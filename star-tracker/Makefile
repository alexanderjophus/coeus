SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)
VER := 0.14

.PHONY: all
all : test build docker

.PHONY: build
build : bin/star-tracker bin/star-tracker-pachd

.PHONY: docker
docker : bin/star-tracker-pachd.tar.gz

bin/star-tracker: $(SRC)
	go build -o ./bin/star-tracker .

bin/star-tracker-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/statsapi/api/v1/game/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/star-tracker-pachd .
	# CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" -installsuffix cgo -o ./bin/star-tracker-pachd .
	upx --brute bin/star-tracker-pachd

bin/star-tracker-pachd.tar.gz: bin/star-tracker-pachd star-tracker.json.tmpl
	./build_save.sh

pipeline : star-tracker.json
	pachctl delete-pipeline star-tracker
	pachctl create-pipeline -f star-tracker.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf bin/
	rm -rf test
	rm -rf out/
