NAME = terraform-provider-convox
VERSION := $(shell git describe --tags --always --dirty)
OS := $(shell go env GOOS)
XC_OS := darwin linux
bindir ?= ~/.terraform.d/plugins

all: tools build

build: $(OS)

build-all: $(XC_OS)

$(XC_OS):
	GOOS=$@ go build -o build/$@/$(NAME) -ldflags "-X main.version=$(VERSION)" main.go

install: build
	install -m 755 build/$(OS)/$(NAME) $(bindir)/$(NAME)

deps:
	glide install

clean:
	-rm -rf build/*

tools:
	curl -sSL https://glide.sh/get | sh


.PHONY: all build install clean
