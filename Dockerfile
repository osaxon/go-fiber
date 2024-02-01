FROM golang:1.21@sha256:cffaba795c36f07e372c7191b35ceaae114d74c31c3763d442982e3a4df3b39e

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Install any dependencies you might need for your build process
RUN apk add --no-cache git

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download

# Copy the entire application source code to the container
COPY . .

# Build the Go application
RUN go build -o out cmd/main.go

# Create a minimal runtime image
FROM alpine:3.14

# Set the working directory in the runtime image
WORKDIR /app

# Copy the binary from the builder image to the runtime image
COPY --from=builder /app/app .

# Expose the port the application will run on
EXPOSE 3000

# Command to run the executable
CMD ["./app"]
