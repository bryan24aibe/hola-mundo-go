# Go Hello World Application

This is a simple "Hello World" application built using Go programming language. The project demonstrates a basic HTTP server that returns a "Hello, World!" message when accessed via a web browser.

## Features

- Simple HTTP server built using Go.
- Responds with "Hello, World!" at the root endpoint (`/`).
- Dockerized version for easy deployment.

## Requirements

To run this project locally, you need to have the following installed:

- [Go (Golang)](https://golang.org/doc/install) version 1.20 or later
- [Docker](https://www.docker.com/get-started) (if you want to run the project inside a container)

### Optional: Docker for Containerization

- If you want to run the project inside a Docker container, Docker must be installed.

## Installation

### 1. Clone the repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/bryan24aibe/hola-mundo-go
cd go-hello-world
```

### 2. Running the Application Locally (Without Docker)

To run the application locally without Docker, follow these steps:

1. **Install Go** if you haven't already by following the [Go installation guide](https://golang.org/doc/install).

2. **Build and run the Go application**:

   In the project directory, run the following command to start the server:

   ```bash
   go run main.go
   ```

3. Open your browser and navigate to [http://localhost:8080](http://localhost:8080). You should see the "Hello, World!" message.

### 3. Running the Application with Docker

If you want to run the application in a Docker container, follow these steps:

1. **Build the Docker image**:

   First, build the Docker image using the following command:

   ```bash
   docker build -t go-hello-world .
   ```

2. **Run the Docker container**:

   After building the image, run it in a container:

   ```bash
   docker run -p 8080:8080 go-hello-world
   ```

3. Open your browser and navigate to [http://localhost:8080](http://localhost:8080). You should see the "Hello, World!" message.

## Code Structure

- `main.go`: The main file containing the Go code that creates the HTTP server and handles the root route (`/`).
- `Dockerfile`: The file used to containerize the Go application.

## Dockerfile

The `Dockerfile` is used to create a Docker image for the Go application. It uses a lightweight Go base image, builds the application, and exposes the required port (8080).

```dockerfile
# Use the official Go base image
FROM golang:1.20-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the app
CMD ["./main"]
```

## Usage

- **HTTP Server**: The application runs a simple HTTP server on port 8080. When you visit `http://localhost:8080`, it responds with "Hello, World!".

## Contributing

Contributions are welcome! If you'd like to help improve this project, feel free to fork the repository, create a new branch, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
