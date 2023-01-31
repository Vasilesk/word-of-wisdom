lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -count=1 ./... -covermode=atomic -v -race

build-server:
	go build -o ./bin/server ./cmd/server

build-client:
	go build -o ./bin/client ./cmd/client

build-all: build-server build-client

docker-build-server:
	docker build -t word-of-wisdom/server -f ./cmd/server/Dockerfile .

docker-build-client:
	docker build -t word-of-wisdom/client -f ./cmd/client/Dockerfile .

docker-build-all: docker-build-server docker-build-client

.PHONY: lint test
