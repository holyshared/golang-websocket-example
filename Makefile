.PHONY: setup

setup:
	go get golang.org/x/net/websocket
	go get github.com/Sirupsen/logrus

.PHONY: build

build:
	go build -o server

.PHONY: clean

clean:
	rm server
