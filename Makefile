all: files

files: generate
	CGO_ENABLED='0' GOOS='linux' GOARCH='amd64' go build -a -ldflags '-extldflags "-static"' -o files files.go

docker: files generate
	docker build -t sparkymat/files .

generate:
	go generate ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./... --timeout=2m

test:
	go test ./...

.PHONY: files
