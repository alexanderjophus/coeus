SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)

.PHONY: all
all : test build docker

.PHONY: build
build : bin/play-tracker bin/play-tracker-pachd

.PHONY: docker
docker : bin/play-tracker-pachd.tar.gz

bin/play-tracker: $(SRC)
	go build -o ./bin/play-tracker .

bin/play-tracker-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/statsapi/api/v1/game/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/play-tracker-pachd .
	upx --brute bin/play-tracker-pachd

bin/play-tracker-pachd.tar.gz: bin/play-tracker-pachd play-tracker.json.tmpl
	./build_save.sh

pipeline : play-tracker.json
	pachctl delete-pipeline play-tracker
	pachctl create-pipeline -f play-tracker.json

test : $(SRC)
	go test -v ./... | tee test

.PHONY : clean
clean: 
	rm -rf bin/
	rm -rf test
	rm -rf out/
	rm play-tracker.json
