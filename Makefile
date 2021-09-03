export GOPROXY=https://proxy.golang.org

SHELL= /bin/bash
GO ?= go
BUILD_DIR := ./bin
BIN_DIR := /usr/local/bin
NAME := psgo
BATS_TESTS := *.bats

# Not all platforms support -buildmode=pie
ifeq ($(shell $(GO) env GOOS),linux)
	ifeq (,$(filter $(shell $(GO) env GOARCH),mips mipsle mips64 mips64le ppc64))
		GO_BUILDMODE := "-buildmode=pie"
	endif
endif
GO_BUILD := $(GO) build $(GO_BUILDMODE)

all: validate build

.PHONY: build
build:
	 $(GO_BUILD) -o $(BUILD_DIR)/$(NAME) ./sample

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: validate
validate:
	golangci-lint run

.PHONY: test
test: test-unit test-integration

.PHONY: test-integration
test-integration:
	bats test/$(BATS_TESTS)

.PHONY: test-unit
test-unit:
	go test -v ./...

.PHONY: install
install:
	sudo install -D -m755 $(BUILD_DIR)/$(NAME) $(BIN_DIR)

.PHONY: uninstall
uninstall:
	sudo rm $(BIN_DIR)/$(NAME)
