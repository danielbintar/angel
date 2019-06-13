include .env
export

test:
	go run test-script/mysql-migration/main.go
	go test -coverprofile ./test-cov `go list ./... | grep -v test-script | grep -v migration`
