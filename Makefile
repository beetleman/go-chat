.PHONY: clean test all

client:
	go build ./cmd/client

server:
	go build ./cmd/server

all: test client server

test:
	go test ./internal/...

clean:
	rm -f client
	rm -f server
