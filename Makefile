MAIN=./onion/cmd/main.go

build:
	-go build $(MAIN)

deps:
	-dep ensure
