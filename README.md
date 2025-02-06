# Rate Limiter

Rate Limiter is a simple and effective API rate limiting solution built with Go, Fiber, and Redis. This project includes an IP-based request counter middleware, Redis integration for managing counters and expirations, and a custom logger to track events.

## Features

- **Sliding Window Rate Limiting:** Applies a request limit per IP or API key within a defined time window.
- **Custom Logger:** Logs events in JSON format to both the database and log files.
- **Redis Integration:** Manages request counters and expiration times using Redis.
- **Fiber Web Framework:** Built on the high-performance and easy-to-use Fiber framework.
- **API Key Support:** Implements API key-based authentication to protect endpoints.

## Project Architecture

- **`cmd/app`**: The entry point of the application where the Fiber web server is initialized and middleware is set up.
- **`internal/logger`**: Contains custom logger functions.
- **`internal/redis`**: Manages Redis operations such as incrementing counters and setting expiration times.
- **`internal/database`**: Handles SQLite database connections for logging API events.
- **`pkg/middleware`**: Contains rate limiting and API key authentication middleware.
- **`pkg/handlers`**: Manages API request handlers.

## Requirements

- [Go](https://golang.org) 1.18 or higher
- [Docker](https://www.docker.com) (recommended for running Redis and containerized environments)
- [Redis](https://redis.io) (for local development or via Docker)

## Getting Started

### Running Locally

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/AliRizaAynaci/rate-limiter.git
   cd rate-limiter
   ```

2. **Run Redis (Using Docker):**

   ```bash
   docker run -d --name redis -p 6379:6379 redis:alpine
   ```
   Alternatively, you can use the provided Docker Compose file.

3. **Set Environment Variables:**

   Optionally, create a `.env` file in the project root:

   ```dotenv
   REDIS_ADDR=localhost:6379
   LOG_DIR=logs
   ```

4. **Build and Run the Application:**

   ```bash
   make build
   make run
   ```
   By default, the server will run on port 3000.

5. **Test the API:**

   Use an API client like Insomnia or Postman and send a request to:

   ```http
   http://localhost:3000/
   ```

### Running with Docker Compose

The project comes with a `docker-compose.yml` file to run both Redis and the application in containers.

1. **Start the Containers:**

   ```bash
   make up
   ```

2. **Access the Application:**

   Visit:

   ```http
   http://localhost:3000/
   ```

3. **View Logs:**

   ```bash
   make logs
   ```

4. **Stop Containers:**

   ```bash
   make down
   ```

## Logging Format

The logger records events in **JSON format**. Example log entry:

```json
{
  "level": "INFO",
  "timestamp": "2024-02-07T12:34:56Z",
  "message": "Rate limit check passed: 192.168.1.1",
  "endpoint": "/api/protected-endpoint"
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|---------|-------------|
| GET    | `/` | Check if the API is running |
| GET    | `/logs` | Retrieve log entries in JSON format |
| GET    | `/api/protected-endpoint` | Access an endpoint protected with an API key |

## Project Structure

```bash
rate-limiter/
├── cmd/
│   └── app/
│       └── main.go           # Entry point of the application
├── internal/
│   ├── logger/
│   │   └── logger.go         # Custom logger implementation
│   ├── redis/
│   │   └── redis_client.go   # Redis client and related functions
│   ├── database/
│   │   └── database.go       # SQLite database connection
├── pkg/
│   ├── middleware/
│   │   ├── rate_limiter.go   # Sliding Window Rate Limiting middleware
│   │   ├── api_key.go        # API Key authentication middleware
│   ├── handlers/
│   │   └── logs_handler.go   # Handler to retrieve logs
├── logs/                     # Directory for log files (mounted via Docker Compose)
├── docker-compose.yml        # Docker Compose configuration
├── Makefile                  # Makefile for running tasks
├── README.md                 # This file
└── go.mod                    # Go module definition
```
