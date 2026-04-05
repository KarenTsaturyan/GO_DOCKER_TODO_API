# Multi-stage build {May be improved}
# The first stage builds the Go application using the official Golang image.
# CGO_ENABLED=0 disables CGO, which allows the binary to be statically linked and run in a minimal container without needing libc.
FROM golang:1.26-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/todoapp/exe /app/cmd/todoapp

# The second stage uses a minimal Alpine image to run the compiled binary.
FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/todoapp/exe /app
CMD ["/app/exe"]