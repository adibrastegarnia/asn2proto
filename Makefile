export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

all:build

build:
	go build  -o build/_output/parser_example  ./cmd/parser_example
	go build  -o build/_output/listener_example  ./cmd/listener_example


clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output

