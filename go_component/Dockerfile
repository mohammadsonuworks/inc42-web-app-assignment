# First stage: Build the Go application
FROM golang:1.22 as builder

# Set the working directory for the build stage
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go binary
RUN go build -o myapp main.go  # Specifying the main file to build

# Second stage: Create a minimal execution environment
FROM alpine:latest  

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled binary from the builder stage to the runtime environment
COPY --from=builder /app/myapp .

# Set the entrypoint to run the Go binary
ENTRYPOINT ["/app/myapp"]
