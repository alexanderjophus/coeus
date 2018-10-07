all: build

build: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/star-tracker .

docker-build: build ## some steps

pipeline:
	pachctl create-pipeline -f star-tracker.json

test: ## runs tests
	go test -v ./...
