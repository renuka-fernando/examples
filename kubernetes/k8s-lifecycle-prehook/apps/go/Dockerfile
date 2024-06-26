# Use the official Golang image to create a build artifact.
FROM golang:1.22.2 as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy the Go Modules manifests.
# COPY go.mod go.sum ./
COPY go.mod ./

# Download the dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 go build -o /go/bin/go-app .

# Use a minimal Docker image to run the Go app.
FROM alpine:3.19.1

# Set the current working directory inside the container.
WORKDIR /app

# Copy the binary from the builder stage.
COPY --from=builder /go/bin/go-app .

# Expose the port on which the app will run.
EXPOSE 8080

# Command to run the binary.
ENTRYPOINT ["./go-app"]
# CMD /app/go-app -port 8080
