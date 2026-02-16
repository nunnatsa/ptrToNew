BUMP ?= patch

.PHONY: build
build:
	go build -o bin/ptrToNew ./cmd/ptr_to_new

.PHONY: build-all
build-all:
	GOOS=linux   GOARCH=amd64 go build -o bin/ptrToNew-linux-amd64       ./cmd/ptr_to_new
	GOOS=linux   GOARCH=arm64 go build -o bin/ptrToNew-linux-arm64       ./cmd/ptr_to_new
	GOOS=darwin  GOARCH=amd64 go build -o bin/ptrToNew-darwin-amd64      ./cmd/ptr_to_new
	GOOS=darwin  GOARCH=arm64 go build -o bin/ptrToNew-darwin-arm64      ./cmd/ptr_to_new
	GOOS=windows GOARCH=amd64 go build -o bin/ptrToNew-windows-amd64.exe ./cmd/ptr_to_new
	GOOS=windows GOARCH=arm64 go build -o bin/ptrToNew-windows-arm64.exe ./cmd/ptr_to_new

.PHONY: test
test:
	go test ./...

.PHONY: release
release:
	./hack/release.sh $(BUMP)

