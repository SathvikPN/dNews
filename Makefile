# ====================================
# Makefile for building dNews service
# ====================================

GO_BUILD = go build
GO_RUN = go run 

# distributable artifacts directory
DIST_DIR = ./dist
CMD_DIR = ./cmd
INTERNAL_DIR = ./internal 

VERSION = "v0.0.0"
LDFLAGS = -X main.Version=${VERSION}

## help: helps with listing available commands
.PHONY: help
help:
	@echo "build: builds the project"
	@echo "clean: cleans up built code and vendor directory"

## lint: runs linters 
.PHONY: lint 
lint:
	@golangci-lint run 

## build: builds the project
.PHONY: build 
build: 
	@mkdir -p dist
	${GO_BUILD} -ldflags "${LDFLAGS}" -o dist/ ./cmd/...

## clean: cleans up built code and vendor directory
.PHONY: clean
clean:
	rm -rf ${DIST_DIR} vendor

.PHONY: run 
run:
	${GO_RUN} ${CMD_DIR}/dnews-rest/main.go