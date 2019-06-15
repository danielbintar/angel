include .env
export $(shell sed 's/=.*//' .env)

test:
	go run test-script/mysql-migration/main.go
	go test -coverprofile ./test-cov `go list ./... | grep -v test-script | grep -v migration | grep -v factory`
