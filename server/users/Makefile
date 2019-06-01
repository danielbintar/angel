test:
	./script/migrate-test
	ENVIRONMENT=test go test -coverprofile ./test-cov `go list ./... | grep -v cmd | grep -v migration | grep -v factory`
