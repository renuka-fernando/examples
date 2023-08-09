FROM --platform=linux/arm64 alpine:3.18.2

USER 10014

COPY go-app-arm64 /bin/go-app
EXPOSE 8080
ENTRYPOINT ["/bin/go-app"]
