MAIN=./cmd/onion/main.go

build:
	-go build $(MAIN)

deps:
	-dep ensure
