SHELL=/bin/bash

OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

.PHONY: up
up:
	cd test && docker compose up -d

.PHONY: down
down:
	cd test && docker compose down -v

.PHONY: go-test
go-test:
	cd test && go test -v ./...

.PHONY: log
log:
	docker logs hackathon_test

.PHONY: test
test:
	-$(MAKE) up
	sleep 15
	-$(MAKE) go-test
	-$(MAKE) down

