CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o Jocy-linux-arm64 -ldflags '-s -w -extldflags "-static"'
upx Jocy-*