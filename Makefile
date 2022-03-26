all: files

files:
	CGO_ENABLED='0' GOOS='linux' GOARCH='amd64' go build -a -ldflags '-extldflags "-static"' -o files files.go

docker: files
	docker build -t sparkymat/files .

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./... --timeout=2m

test:
	go test ./...

.PHONY: files
