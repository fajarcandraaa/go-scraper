tidy:
	@go mod tidy
	@go mod vendor

clean:
	@go clean -cache

start:
	@go run main.go

update:
	@go get -u