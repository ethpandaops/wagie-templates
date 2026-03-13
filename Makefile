.DEFAULT_GOAL := help
.PHONY: test tidy clean help

WAGIE_CORE_DIR ?= ../wagie

## test: validate this library together with Wagie core templates
test:
	WAGIE_CORE_DIR="$(WAGIE_CORE_DIR)" go test -shuffle=on ./...

## tidy: tidy go modules
tidy:
	go mod tidy

## clean: remove Go test cache for this module
clean:
	go clean -testcache

## help: show this help
help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
