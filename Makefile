GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
PROGRAM := mcping
ifeq ($(GOOS),windows)
	PROGRAM := $(PROGRAM).exe
	GOFLAGS := -v -buildmode=exe
else
	GOFLAGS := -v
endif

CGO_ENABLED := 0
RELEASE_VERSION := $(shell git describe --tags --always || echo "local")

all: test build

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -ldflags "-X main.version=$(RELEASE_VERSION)" -o $(PROGRAM)

.PHONY: test
test:
	go test -count=1 -v -p 1 ./...

.PHONY: clean
clean:
	rm -f $(PROGRAM)