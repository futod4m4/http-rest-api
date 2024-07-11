.PHONY: build
build:
		$(SET_ENV) go build ./cmd/apiserver

.PHONY: test
test:
	go test -v -timeout 30s ./ ...

.DEFAULT_GOAL := build