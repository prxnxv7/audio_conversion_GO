# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Install dependencies
RUN apk add --no-cache ffmpeg

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o audio_conversion main.go

# Stage 2: Run the application in a minimal environment
FROM alpine:latest

# Install ffmpeg in the runtime image
RUN apk add --no-cache ffmpeg

# Copy the built binary from the builder stage
COPY --from=builder /app/audio_conversion /audio_conversion

# Expose the port your application uses
EXPOSE 8080

# Run the application
CMD ["/audio_conversion"]
