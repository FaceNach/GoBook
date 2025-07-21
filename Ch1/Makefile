BINARY_NAME=hello_world.exe

.DEFAULT_GOAL := build


.PHONY:fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
    del :$(BINARY_NAME)