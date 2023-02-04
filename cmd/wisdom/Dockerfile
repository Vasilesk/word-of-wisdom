FROM golang:1.19.3-alpine3.17 as build

COPY . /build

WORKDIR /build

RUN go build -o /app ./cmd/wisdom

FROM alpine:3.17

COPY --from=build /app /app

ENTRYPOINT /app
