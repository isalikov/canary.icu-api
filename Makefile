default: build

.PHONY: build
build:
	go build -o ./target/api

.PHONY: lint
lint:
	golint
