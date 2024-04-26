#!/bin/bash
rm bin/gSandpiles-*
env GOES=linux GOARCH=amd64 go build -o bin/gSandpiles-linux-amd64
# env GOOS=linux GOARCH=arm64 go build -o bin/gSandpiles-linux-arm64
# env GOOS=linux GOARCH=arm go build -o bin/gSandpiles-linux-arm
# env GOOS=netbsd GOARCH=amd64 go build -o bin/gSandpiles-netbsd-amd64
# env GOOS=netbsd GOARCH=arm64 go build -o bin/gSandpiles-netbsd-arm64
# env GOOS=openbsd GOARCH=amd64 go build -o bin/gSandpiles-openbsd-amd64
# env GOOS=openbsd GOARCH=arm64 go build -o bin/gSandpiles-openbsd-arm64
