test:
	ENVIRONMENT=test go test -coverprofile ./test-cov `go list ./... | grep -v cmd | grep -v model`

pretty:
	gofmt -s -w .

build:
	go build cmd/main.go

run: build
	./main
