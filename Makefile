SHELL=/bin/bash

OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

.PHONY: up
up:
	docker compose up -d

.PHONY: down
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
	-$(MAKE) up
	sleep 15
	-$(MAKE) go-test
	-$(MAKE) down

.PHONY: dev-up
dev-up:
	docker compose --file docker-compose.dev.yml up -d
	sleep 30
	go run ./cmd/main_dev.go

.PHONY: dev-down
dev-down:
	docker compose --file docker-compose.dev.yml down -v

.PHONY: add-to-push
add-to-push:
	git add .
	git commit -m "push"
	git push git@github.com:tetsushi-yamada/hackathon_backend.git