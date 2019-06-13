include .env
export $(shell sed 's/=.*//' .env)

test:
	ENVIRONMENT=test go test -v -coverprofile ./test-cov ./...
