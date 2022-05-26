all: docker

files-linux-amd64: generate
	CGO_ENABLED='0' GOOS='linux' GOARCH='amd64' go build -a -ldflags '-extldflags "-static"' -o files-linux-amd64 files.go

files-linux-arm: generate
	CGO_ENABLED='0' GOOS='linux' GOARCH='arm' go build -a -ldflags '-extldflags "-static"' -o files-linux-arm files.go

files-linux-arm64: generate
	CGO_ENABLED='0' GOOS='linux' GOARCH='arm64' go build -a -ldflags '-extldflags "-static"' -o files-linux-arm64 files.go

docker: files generate
	docker build -t sparkymat/files .

generate:
	go generate ./...

lint: generate
	golangci-lint run

test: generate
	go test ./...

.PHONY: files
