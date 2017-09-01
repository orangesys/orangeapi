NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

IMPORT_PATH := github.com/orangesys/orangeapi
PKG_SRC := $(IMPORT_PATH)/cmd/orangeapi

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := /vendor/

.PHONY: all clean deps build 

BINARYDIR := bin
BINARY := orangeapi
LINUX_AMD64_SUFFIX := _linux-amd64

SOURCEDIR := .
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' -type f)

LDFLAGS := -ldflags="-w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -X \"main.BuildTime=$(BUILDTIME)\" -X \"main.GoVersion=$(GOVERSION)\""

GLIDE := glide
GLIDE_VERSION := 0.12.3

DOCKER_IMAGE_NAME := orangesys/orangeapi
DOCKER_IMAGE_TAG := $(VERSION)
DOCKER_IMAGE := $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.DEFAULT_GOAL := $(BINARYDIR)/$(BINARY)

all: clean deps build

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@/bin/sh -c "PKG_SRC=$(PKG_SRC) VERSION=${VERSION} ./scripts/build.sh"

clean:
	rm -fr $(BINARYDIR)

deps: 
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -u github.com/golang/dep/cmd/dep
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/DATA-DOG/godog/cmd/godog
	@dep ensure

test:
	@test -z "$(gofmt -s -l ./pkg | tee /dev/stderr)"
	@test -z "$(go vet ./pkg/... | tee /dev/stderr)"
	@go test -v ./pkg/...