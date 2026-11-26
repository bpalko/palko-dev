APP_NAME := ssh-portfolio
BIN_DIR := bin
HOST_KEY := .ssh/host_ed25519

.PHONY: build run test keys fmt tidy clean build-linux

build: | $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/ssh-portfolio

run:
	go run ./cmd/ssh-portfolio

test:
	go test ./...

keys:
	mkdir -p .ssh
	ssh-keygen -t ed25519 -f $(HOST_KEY) -N "" -q
	@echo "Generated host key at $(HOST_KEY)"

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)

build-linux: | $(BIN_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-linux ./cmd/ssh-portfolio

$(BIN_DIR):
	mkdir -p $(BIN_DIR)
