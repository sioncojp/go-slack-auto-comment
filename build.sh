#!/bin/bash -
declare -r Name="go-slack-auto-comment"

for GOOS in darwin linux; do
    GO111MODULE=on GOOS=$GOOS GOARCH=amd64 go build -o bin/go-slack-auto-comment-$GOOS-amd64 *.go
done
