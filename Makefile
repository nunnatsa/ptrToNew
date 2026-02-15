.PHONY: build
build:
	go build -o bin/ptrToNew ./cmd/ptr_to_new

.PHONY: test
test:
	go test ./...

