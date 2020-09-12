run:
	@go run main.go

generate:
	go generate ./...

fmt:
	go fmt ./...

ver:
	go vet ./...
