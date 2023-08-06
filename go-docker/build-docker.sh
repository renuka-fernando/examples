#!/bin/bash

IMAGE_NAME=renukafernando/go-docker:v1

rm go-app
env GOOS=linux GOARCH=amd64 go build -o go-app .
# env GOOS=linux GOARCH=arm64 go build -o request-info-arm64 .

#docker buildx build --push --platform linux/arm64/v8,linux/amd64 -t "$IMAGE_NAME" .
docker build --build-arg TARGETARCH=amd64 -t "$IMAGE_NAME" .
