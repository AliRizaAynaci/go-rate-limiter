# Rate Limiter

Rate Limiter is a simple and effective API rate limiting solution built with Go, Fiber, and Redis. This project includes an IP-based request counter middleware, Redis integration for managing counters and expirations, and a custom logger to track events.

## Features

- **Rate Limiting:** Applies a request limit for each IP address.
- **Custom Logger:** Logs rate limiter events to a log file.
- **Redis Integration:** Uses Redis to manage counters and expiration settings.
- **Fiber Web Framework:** Built on the high-performance and easy-to-use Fiber framework.

## Project Architecture

- **`cmd/app`**: The entry point of the application. This is where the Fiber web server is initialized and the middleware is set up.
- **`internal/logger`**: Contains custom logger functions.
- **`internal/redis`**: Handles Redis operations such as incrementing counters and setting expiration.
- **`pkg/middleware`**: Contains the rate limiting middleware.

## Requirements

- [Go](https://golang.org) 1.18 or higher
- [Docker](https://www.docker.com) (recommended for running Redis and containerized environments)
- [Redis](https://redis.io) (for local development or via Docker)

## Getting Started

### Running Locally

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/rate-limiter.git
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

4. **Run the Application:**

   ```bash
   go run cmd/app/main.go
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
   docker-compose up --build
   ```

2. **Access the Application:**

   Visit:

   ```http
   http://localhost:3000/
   ```

3. **Logs:**

   Application logs will be mapped to the `./logs` directory on your host machine.

## Tests and Coverage

This project includes unit tests. To run all tests and view the coverage summary, execute:

```bash
go test -cover ./...
```

For a detailed HTML coverage report, run:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Project Structure

```bash
rate-limiter/
├── cmd/
│   └── app/
│       └── main.go           # Entry point of the application
├── internal/
│   ├── logger/
│   │   └── logger.go         # Custom logger implementation
│   └── redis/
│       └── redis.go          # Redis client and related functions
├── pkg/
│   └── middleware/
│       └── rate_limiter.go   # Rate limiting middleware
├── logs/                     # Directory for log files (mounted via Docker Compose)
├── docker-compose.yml        # Docker Compose configuration
├── README.md                 # This file
└── go.mod                    # Go module definition
```

