#!/bin/sh

# Build the binary
GOOS=linux GOARCH=arm go build -ldflags="-s -w" -trimpath -o target/alert ./cmd/alert/main.go

# Check for --install argument
if [ "$1" = "--install" ]; then
    INSTALL_DIR="${PREFIX:-/usr/bin}"
    echo "Installing alert to $INSTALL_DIR"
    install -m 755 target/alert "$INSTALL_DIR/alert"
fi
