#!/bin/bash

IMAGE_NAME_AMD=renukafernando/go-app:v1-amd64
IMAGE_NAME_ARM=renukafernando/go-app:v1-arm64

rm -f go-app-*
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-app-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o go-app-arm64 .
chmod +x go-app-amd64
chmod +x go-app-arm64

#docker buildx build --push --platform linux/arm64/v8,linux/amd64 -t "$IMAGE_NAME" .
docker build --build-arg TARGETARCH=amd64 -t "$IMAGE_NAME_AMD" .
docker build --build-arg TARGETARCH=arm64 -t "$IMAGE_NAME_ARM" .
