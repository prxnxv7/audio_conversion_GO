# WAV to FLAC Audio Conversion Service

## Overview

This project is a robust backend service built in Go that converts WAV audio streams to FLAC format in real-time. The service utilizes WebSockets for real-time streaming of converted audio data back to the client. It is designed to efficiently handle multiple simultaneous audio streams while ensuring minimal latency and high fidelity in the conversion process.

## Features

- Real-time conversion of WAV audio streams to FLAC format.
- Streaming of converted FLAC data back to the client via WebSockets.
- API documentation using Swagger.
- Efficient handling of concurrent audio streams.
- Error handling for streaming and conversion processes.

## Technologies Used

- Go (Golang)
- Gin (Web Framework)
- WebSockets
- Swagger (API Documentation)
- FFmpeg (for audio processing)


## Setup Instructions

### Prerequisites

- Go 1.18 or higher
- Docker (optional, for containerization)
- FFmpeg (ensure it is installed and accessible in your system's PATH)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/audio_conversion.git
   cd audio_conversion

2. Install the required Go modules:

    go mod tidy

### Running the Application

    go run main.go

The server will start on http://localhost:8080.

- Accessing Swagger Documentation
    Open your browser and navigate to http://localhost:8080/swagger/index.html to view the API documentation.

### API Endpoints
    WebSocket Connection
    Endpoint: ws://localhost:8080/ws
    Method: GET
    Description: Establish a WebSocket connection to send WAV audio data and receive the converted FLAC audio data in real-time.