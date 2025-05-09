# Use a smaller base image for Go
FROM golang:1.24.3-alpine3.21

# Install bash in a single layer to reduce image size
RUN apk add --no-cache bash

# Create a non-root user and set up home directory
RUN adduser -D -h /home/cloak cloak

# Set working directory
WORKDIR /home/cloak/cloak

# Copy application code
COPY . .

# Set ownership and permissions in a single command
RUN chown -R cloak:cloak . && chmod -R 755 .

# Set environment variables
ENV SHELL=/bin/bash \
    XDG_CONFIG_DIRS=/home/cloak/.config

# Initialize configuration as part of the build
RUN go run main.go config init env

# Switch to non-root user
USER cloak

# Default command (optional, add if needed)
# CMD ["go", "run", "main.go"]
