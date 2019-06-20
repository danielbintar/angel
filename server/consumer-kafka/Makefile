test:
	go test -coverprofile ./test-cov `go list ./... | grep -v cmd | grep -v model`

build:
	go build cmd/main.go

run: build
	./main
