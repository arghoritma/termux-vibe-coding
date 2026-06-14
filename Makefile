BINARY := termux-vibe-coding
VERSION := 0.1.0
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")

BUILD_FLAGS := -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT)"

.PHONY: all build clean npm-release

all: build

build:
	go build $(BUILD_FLAGS) -o $(BINARY) ./cmd/termux-vibe-coding

build-android:
	GOOS=android GOARCH=arm64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o $(BINARY)-android-arm64 ./cmd/termux-vibe-coding

build-all:
	GOOS=android GOARCH=arm64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o release/termux-vibe-coding-android-arm64 ./cmd/termux-vibe-coding
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o release/termux-vibe-coding-linux-amd64 ./cmd/termux-vibe-coding
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o release/termux-vibe-coding-darwin-amd64 ./cmd/termux-vibe-coding
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o release/termux-vibe-coding-darwin-arm64 ./cmd/termux-vibe-coding
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o release/termux-vibe-coding-windows-amd64.exe ./cmd/termux-vibe-coding

npm-release: build-all
	cp release/termux-vibe-coding-* npm/bin/
	cd npm && npm pack

clean:
	rm -f $(BINARY) $(BINARY)-android-arm64
	rm -rf release npm/bin/*.tar.gz

test:
	go vet ./...
