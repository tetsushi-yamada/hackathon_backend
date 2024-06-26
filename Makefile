SHELL=/bin/bash

OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

.PHONY: test-up
up:
	docker compose up -d

.PHONY: test-down
down:
	docker compose down -v

.PHONY: go-test
go-test:
	cd test && go test -v ./...

.PHONY: log
log:
	docker logs hackathon_test

.PHONY: test
test:
	-$(MAKE) test-up
	sleep 20
	-$(MAKE) go-test
	-$(MAKE) test-down

.PHONY: dev-up
dev-up:
	docker compose --file docker-compose.dev.yml up -d
	sleep 30
	go run ./cmd/main_dev.go

.PHONY: dev-down
dev-down:
	docker compose --file docker-compose.dev.yml down -v