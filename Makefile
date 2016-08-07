.PHONY: setup

setup:
	go get golang.org/x/net/websocket

.PHONY: build

build:
	go build server.go

.PHONY: clean

clean:
	rm server
