lint:
	golangci-lint run -c .golangci.yml ./...

generate:
	go generate ./...

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

install-mockery:
	go install github.com/vektra/mockery/v2@v2.20.0

.PHONY: lint test
