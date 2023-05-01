# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help build
.DEFAULT_GOAL := help

build: ## builds the project binary
	@export GOARCH=arm64 && go build \
		-tags release \
		-o bin/vehicle cmd/vehicle/main.go

run: ## run binary file
	-bin/vehicle

run-local: ## db and app local run
	-docker-compose up --build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'