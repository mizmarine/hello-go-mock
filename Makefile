run:
	@go run main.go

generate:
	go generate ./...

test:
	go test -v ./...

fmt:
	go fmt ./...

ver:
	go vet ./...
