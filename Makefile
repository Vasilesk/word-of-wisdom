lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -count=1 ./... -covermode=atomic -v -race

build-wisdom:
	go build -o ./bin/wisdom ./cmd/wisdom

build-client:
	go build -o ./bin/client ./cmd/client

build-all: build-wisdom build-client

docker-build-wisdom:
	docker build -t word-of-wisdom/wisdom -f ./cmd/wisdom/Dockerfile .

docker-build-client:
	docker build -t word-of-wisdom/client -f ./cmd/client/Dockerfile .

docker-build-all: docker-build-wisdom docker-build-client

.PHONY: lint test
