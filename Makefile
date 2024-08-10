client:
	go build ./cmd/client

server:
	go build ./cmd/server

all: client server

.PHONY: clean
clean:
	rm -f client
	rm -f server
