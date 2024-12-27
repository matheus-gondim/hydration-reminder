.PHONY: all build clean
.PHONY: all vendor clean

ensure-dependencies:
	go mod tidy

vendor:
	go mod vendor

build: vendor
ifeq (${OS},(Windows_NT))
	if (!(Test-Path ./build)) { mkdir build }
else
	mkdir -p build
endif
	go build -mod=vendor -v -o ./build ./...
