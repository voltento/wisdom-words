# Builder Stage
FROM golang:1.17 as builder

# Set working directory inside the container
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY .. .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Final Stage
FROM alpine:latest

# Copy the server binary from the builder stage
COPY --from=builder /app/server /app/server

# Run the server
ENTRYPOINT ["/app/server", "--difficulty=5"]
