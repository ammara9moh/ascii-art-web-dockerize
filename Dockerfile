# Dockerfile

# --- Stage 1: Build ---
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for better layer caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o server .

# --- Stage 2: Run ---
FROM alpine:3.22

# # Create a non-root user
# RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/server .

# Copy the banner files, templates, and static files directly from build context
COPY banners ./banners
COPY templates ./templates
COPY static ./static

# # Change ownership to non-root user
# RUN chown -R appuser:appgroup /app

# Add metadata labels
LABEL maintainer="ASSEM"
LABEL version="1.0"
LABEL description="ASCII Art Web Server"

# # Switch to non-root user
# USER appuser

# Expose the port your server listens on
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./server"]