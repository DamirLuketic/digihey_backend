# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help build
.DEFAULT_GOAL := help

build: ## builds the digihey backend binary
	@export GOARCH=amd64 && go build \
		-tags release \
		-o bin/digihey cmd/digihey/main.go

run: ## run digihey backend
	-bin/digihey

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


