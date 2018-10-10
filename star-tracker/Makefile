.PHONY: all
all : build test

.PHONY: build
build : bin/star-tracker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/star-tracker .

.PHONY : clean
clean: 
	rm -rf bin/

docker-build : build ## some steps

pipeline :
	pachctl create-pipeline -f star-tracker.json

test : ## runs tests
	go test -v ./...
