include .env
export $(shell sed 's/=.*//' .env)

pretty:
	gofmt -s -w .

test: pretty
	go run test-script/postgres-migration/main.go
	go test -coverprofile ./test-cov `go list ./... | grep -v test-script | grep -v migration | grep -v factory`
