include .env
export $(shell sed 's/=.*//' .env)

build:
	go build cmd/main.go

run: build
	./main
