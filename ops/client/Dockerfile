# Builder Stage
FROM golang:1.17 as builder

# Set working directory inside the container
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o client ./cmd/client

# Final Stage
FROM alpine:latest

# Copy the client binary from the builder stage
COPY --from=builder /app/client /app/client

# Run the client
ENTRYPOINT ["/app/client"]
