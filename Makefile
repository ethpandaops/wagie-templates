.DEFAULT_GOAL := help
.PHONY: test tidy clean help

## test: validate templates against wagie core
test:
	go test -shuffle=on ./...

## tidy: tidy go modules
tidy:
	go mod tidy

## clean: remove Go test cache for this module
clean:
	go clean -testcache

## help: show this help
help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
