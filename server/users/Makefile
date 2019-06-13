include .env
export $(shell sed 's/=.*//' .env)

build:
	go build cmd/main.go

run: build
	./main

migrate-up:
	./script/migrate-up

migrate-down:
	./script/migrate-down

test:
	./script/migrate-test
	ENVIRONMENT=test go test -coverprofile ./test-cov `go list ./... | grep -v cmd | grep -v migration | grep -v factory`
