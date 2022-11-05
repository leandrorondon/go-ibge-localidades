test:
	@go test -race $(go list ./... | grep -v /examples/)

lint:
	@golangci-lint run
