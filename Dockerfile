# Use the official Go image as a base
FROM golang:1.19-alpine

# Install FFmpeg
RUN apk update && apk add --no-cache ffmpeg

# Set working directory
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o audio_conversion ./cmd/server

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./audio_conversion"]
