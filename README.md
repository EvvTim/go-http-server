# Go HTTP Server Boilerplate

## Overview

This repository provides a boilerplate for setting up a basic HTTP server in Go. It uses the [Chi](https://github.com/go-chi/chi) router for HTTP routing and [CORS](https://github.com/go-chi/cors) middleware for Cross-Origin Resource Sharing. This boilerplate is designed to serve as a starting point for building web applications or APIs in Go, featuring essential functionality like environment variable management, routing, and basic error handling.

## Features

- **Router**: Utilizes Chi for routing HTTP requests.
- **CORS Handling**: Configures CORS to control access from different origins.
- **Environment Configuration**: Loads environment variables from a `.env` file using `godotenv`.
- **Basic Endpoints**: Includes sample endpoints for health checks and error handling.
- **Error Handling**: Implements custom error handling for missing configuration.

## Getting Started

### Prerequisites

- Go 1.18 or later
- `godotenv` for environment variable loading
- `chi` for routing
- `cors` for CORS support

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/go-http-server-boilerplate.git
   cd go-http-server-boilerplate
    ```
2. **Install dependencies:**

   ```bash
   go mod tidy
   ```
3. **Create a `.env` file with the following environment variables:**
   ```env
   PORT=PORT_NUMBER
   ```
4. **Run the server:**

   ```bash
   go run main.go
   ```

   The server will start on the specified port, and you can access it at `http://localhost:PORT_NUMBER`.

### Endpoints

- **GET /v1/health**: Returns a simple health check response.
- **GET /v1/err**: Simulates an error response for testing purposes.

### Configuration

- **Environment Variables**: The server reads environment variables from a `.env` file using `godotenv`. You can specify the port number in the `.env` file.

