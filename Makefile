-include .env
RELEASE ?= 0.0.1
BIN_TEST_API := "./dist/api"

# все таргеты, которые не создают новые файлы, должны быть отмечены как PHONY, иначе make вернет в exitcode ошибку
.PHONY: lint run stop run-test-api save-vendor test integrations-test

build-api:
	go build -v -o $(BIN_TEST_API) -ldflags "$(LDFLAGS)" ./cmd/api

run-test-api: build-api
	$(BIN_TEST_API) -config ./configs/api.json

test:
	go test -timeout=5m ./internal/...

integrations-test:
	go test  ./tests/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

run:
	docker-compose up -d

stop:
	docker-compose down

api-swagger:
	swag init --output docs/api --dir cmd/api --parseDependency true --parseInternal true

save-vendor:
	go mod tidy
	go mod vendor
