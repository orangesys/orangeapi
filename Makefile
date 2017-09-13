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

$(BINARYDIR)/$(GLIDE):
	if [ ! -d $(BINARYDIR) ]; then mkdir $(BINARYDIR); fi
ifeq ($(shell uname),Darwin)
	curl -fL https://github.com/Masterminds/glide/releases/download/$(GLIDE_VERSION)/glide-$(GLIDE_VERSION)-darwin-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./darwin-amd64/glide $(BINARYDIR)/$(GLIDE)
	rm -fr ./darwin-amd64
	rm ./glide.zip
else
	curl -fL https://github.com/Masterminds/glide/releases/download/$(GLIDE_VERSION)/glide-$(GLIDE_VERSION)-linux-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./linux-amd64/glide $(BINARYDIR)/$(GLIDE)
	rm -fr ./linux-amd64
	rm ./glide.zip
endif

all: clean deps build

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@/bin/sh -c "PKG_SRC=$(PKG_SRC) VERSION=${VERSION} ./scripts/build.sh"

clean:
	@rm -fr $(BINARYDIR)

deps: 
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	$(BINARYDIR)/$(GLIDE)
		$(BINARYDIR)/$(GLIDE) install

test:
	@test -z "$(gofmt -s -l ./pkg | tee /dev/stderr)"
	@test -z "$(go vet ./pkg/... | tee /dev/stderr)"
	@go test -v ./pkg/...