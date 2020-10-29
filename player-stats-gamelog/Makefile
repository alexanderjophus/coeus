SHELL := /bin/bash
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
SHA := $(shell git rev-parse --verify HEAD)

.PHONY: all
all : test build docker

.PHONY: build
build : bin/player-stats-gamelog bin/player-stats-gamelog-pachd

.PHONY: docker
docker : bin/player-stats-gamelog-pachd.tar.gz

bin/player-stats-gamelog: $(SRC)
	go build -o ./bin/player-stats-gamelog .

bin/player-stats-gamelog-pachd: $(SRC)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w -X main.inDir=/pfs/statsapi/api/v1/people/ -X main.outDir=/pfs/out/" -installsuffix cgo -o ./bin/player-stats-gamelog-pachd .

bin/player-stats-gamelog-pachd.tar.gz: bin/player-stats-gamelog-pachd player-stats-gamelog.json.tmpl
	./build_save.sh

pipeline : player-stats-gamelog.json
	pachctl delete pipeline player-stats-gamelog
	pachctl create pipeline -f player-stats-gamelog.json

test:
	go test -v ./...

.PHONY : clean
clean: 
	rm -rf bin/
	rm -rf test
	rm -rf out/
