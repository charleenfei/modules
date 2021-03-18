#!/bin/bash

PACKAGES=$(shell go list ./... | grep -v '/simulation')
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=NameService \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=nsd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=nscli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

# all: install

# install: go.sum
# 		go install -mod=readonly $(BUILD_FLAGS) ./cmd/nsd
# 		go install -mod=readonly $(BUILD_FLAGS) ./cmd/nscli

# installWithFaucet: go.sum
# 		go install -mod=readonly $(BUILD_FLAGS) -tags faucet ./cmd/nsd
# 		go install -mod=readonly $(BUILD_FLAGS) -tags faucet ./cmd/nscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen:v0.1 sh ./scripts/protocgen.sh

format-go: ## formats go code
	@golangci-lint run --disable-all --enable=goimports --fix ./...	

lint-check: ## runs linters on go code
	@golangci-lint run --skip-dirs=build/*  --disable-all --enable=golint --enable=goimports --enable=vet --enable=nakedret \
	--enable=unused --skip-dirs=resources --skip-dirs=testingutils --timeout=2m ./...;
