# WAV to FLAC Audio Conversion Service

## Project Overview
A Go-based backend service for real-time conversion of WAV audio streams to FLAC format, with support for WebSocket streaming.

## Prerequisites
- Go 1.19 or later
- FFmpeg (for audio conversion)
- Docker (optional for containerization)

## Project Structure
- `cmd/server`: Entry point for the application.
- `internal`: Contains core logic, such as audio conversion and WebSocket handling.
- `test`: Test files for conversion logic and WebSocket connections.

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/wav-to-flac-converter.git
    cd wav-to-flac-converter
    ```

2. **Install FFmpeg** (if not using Docker):
    ```bash
    sudo apt update && sudo apt install -y ffmpeg
    ```

3. **Run the server**:
    ```bash
    go run cmd/server/main.go
    ```

4. **Run tests**:
    ```bash
    go test ./test/...
    ```

## Usage

- The WebSocket endpoint is available at `ws://localhost:8080/ws`.
- Stream WAV data to the endpoint, and you will receive FLAC data in real time.

## Docker

Build and run the project in Docker:
```bash
docker build -t wav-to-flac-converter .
docker run -p 8080:8080 wav-to-flac-converter
