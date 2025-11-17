# Variables
BINARY_NAME=jetbra-free
BIN_DIR=./bin

all: build

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd

run: build
	./bin/$(BINARY_NAME)

build-all: build-mac build-mac-arm build-windows build-win7 build-linux

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd

build-win7:
	GOOS=windows GOARCH=amd64 go build -ldflags="-extldflags=-subsystem=console,6.1" -o $(BIN_DIR)/$(BINARY_NAME)-windows7-amd64.exe ./cmd

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd

clean:
	rm -rf ./bin/

test-embed:
	go run test_embed.go cmd/assets.go

# Legacy targets (for documentation)
# These are no longer needed with Go embed:
# install-bindata: go install github.com/go-bindata/go-bindata/v3/go-bindata@latest  
# bindata-access: go-bindata -o internal/util/access.go -pkg util static/... templates/... cache/...