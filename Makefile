.PHONY: all build-linux build-linux-arm64 build-linux-amd64 compress clean

# Default target
all: build-linux compress

# Build for all Linux architectures
build-linux: build-linux-arm64 build-linux-amd64

# Build for Linux ARM64
build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o Jocy-linux-arm64 -ldflags '-s -w -extldflags "-static"'

# Build for Linux AMD64
build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Jocy-linux-amd64 -ldflags '-s -w -extldflags "-static"'

# Compress binaries with UPX
compress:
	upx Jocy-*

# Clean build artifacts
clean:
	rm -f Jocy-*

# 完整打包
.PHONY: full
full: build-linux compress