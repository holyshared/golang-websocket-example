.PHONY: setup

setup:
	go get golang.org/x/net/websocket

.PHONY: build

build:
	go build -o server

.PHONY: clean

clean:
	rm server
