export CGO_ENABLED=0
export GO111MODULE=on

ANTLR_VERSION := latest

.PHONY: build

all:build

build:
	go build  -o build/_output/parser_example  ./cmd/parser_example
	go build  -o build/_output/listener_example  ./cmd/listener_example


coverage: # @HELP generate unit test coverage data
coverage: deps build

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"


linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 30m

license_check: # @HELP examine and ensure license headers exist
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

antlr-docker: # @HELP build antlr Docker image
antlr-docker:
	docker build . -f build/antlr/Dockerfile \
		--build-arg ANTLR_VERSION=${ANTLR_VERSION} \
		-t onosproject/antlr:${ANTLR_VERSION}


clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output

