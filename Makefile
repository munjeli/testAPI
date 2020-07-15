all: build test

build:
	go build -o testapi

test:
	./tests.sh

run:
	go run main.go &

clean: SHELL:=/bin/bash
clean:
	kill $$(lsof -i :443 -t) && rm -f slacktest

pack: build
	./package.sh


