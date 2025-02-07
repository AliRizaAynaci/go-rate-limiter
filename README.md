# Rate Limiter

Rate Limiter is an **API rate limiting** solution built with **Go, Fiber, and Redis**. It uses a **sliding window algorithm** integrated with Redis and supports **API key authentication**. Additionally, it provides detailed metrics monitoring with **Prometheus and Grafana**.

## 🚀 Features

- ✅ **Sliding Window Rate Limiting** → Rate limiting based on IP or API key.
- ✅ **Prometheus Metrics** → Monitor API requests, rate limit violations, and Redis usage.
- ✅ **Grafana Dashboard** → Visualize Prometheus data.
- ✅ **Custom Logger** → Store logs in JSON format in a database and log files.
- ✅ **Redis Integration** → Manage counters and expiration times using Redis.
- ✅ **Fiber Framework** → High-performance Go web framework.
- ✅ **API Key Support** → API key authentication mechanism.

## 🏷 Project Architecture

### 📂 Important Directories:
- **`cmd/app`** → Entry point of the application, initializes the Fiber server.
- **`internal/logger`** → Custom logging functions.
- **`internal/redis`** → Redis operations (counter management, expiration setting).
- **`internal/database`** → SQLite database management.
- **`internal/prometheus`** → Prometheus metrics management.
- **`pkg/middleware`** → Rate limiting and API key authentication middleware.
- **`pkg/handlers`** → API request handlers.

## 🔧 Requirements

- [Go](https://golang.org) **1.18+**
- [Docker](https://www.docker.com) **(For Redis, Prometheus, and Grafana)**
- [Redis](https://redis.io)
- [Prometheus](https://prometheus.io)
- [Grafana](https://grafana.com/)

## 🚀 Getting Started

### 🌍 **Local Development**

1️⃣ **Clone the Repository:**
   ```sh
   git clone https://github.com/AliRizaAynaci/go-rate-limiter.git
   cd rate-limiter
   ```

2️⃣ **Run Redis with Docker:**
   ```sh
   docker run -d --name redis -p 6379:6379 redis:alpine
   ```

3️⃣ **Set Up Environment Variables (`.env` file):**
   ```dotenv
   REDIS_ADDR=localhost:6379
   LOG_DIR=logs
   ```

4️⃣ **Run the Go Application:**
   ```sh
   make build
   make run
   ```

5️⃣ **Test the API:**
   ```sh
   curl http://localhost:3000/
   ```

## 📦 Running with Docker Compose

Run all services (Redis, Prometheus, Grafana, and the application) using Docker Compose.

1️⃣ **Start Services:**
   ```sh
   make up
   ```

2️⃣ **Access the API:**
   ```sh
   http://localhost:3000/
   ```

3️⃣ **Access Grafana Dashboard:**
   ```sh
   http://localhost:3001/
   ```

4️⃣ **View Logs:**
   ```sh
   make logs
   ```

5️⃣ **Stop Services:**
   ```sh
   make down
   ```

## 📊 Prometheus & Grafana Integration

### **📌 Viewing Prometheus Metrics**
- **Metrics Endpoint:**  
  ```sh
  http://localhost:3000/metrics
  ```
- **Prometheus UI (when running in Docker):**  
  ```sh
  http://prometheus:9090/
  ```

### **📌 Supported Metrics**
| Metric Name                     | Description |
|----------------------------------|-------------|
| `redis_total_requests`          | Total number of Redis requests |
| `redis_requests_per_endpoint`   | Redis requests per endpoint |
| `redis_rate_limit_violations`   | Number of rate limit violations |

> **📌 PromQL Queries**
```promql
redis_total_requests
```
```promql
redis_requests_per_endpoint
```
```promql
redis_rate_limit_violations
```

## 🌍 API Endpoints

| Method | Endpoint | Description |
|--------|---------|-------------|
| GET    | `/` | Check if the API is running |
| GET    | `/logs` | Retrieve logs in JSON format |
| GET    | `/metrics` | Display Prometheus metrics |
| GET    | `/api/protected-endpoint` | Requires an API key |

## 🏢 Project Structure

```bash
rate-limiter/
├── cmd/
│   └── app/
│       ├── logs.db           # SQLite log database
│       └── main.go           # Main application file
├── internal/
│   ├── database/
│   │   └── database.go       # SQLite connection
│   ├── logger/
│   │   ├── logger.go         # Logger functions
│   │   └── logger_test.go    # Logger unit tests
│   ├── models/
│   │   ├── api_key.go        # API key model
│   │   └── log_entry.go      # Log entry model
│   ├── redis/
│   │   ├── redis_client.go   # Redis client and functions
│   │   └── redis_test.go     # Redis unit tests
│   ├── prometheus/
│   │   └── metrics.go        # Prometheus metrics management
├── pkg/
│   ├── middleware/
│   │   ├── api_key.go        # API Key authentication middleware
│   │   ├── middleware_test.go# Middleware unit tests
│   │   └── rate_limiter.go   # Sliding Window Rate Limiting middleware
│   ├── handlers/
│   │   └── logs_handler.go   # Log retrieval handler
├── logs/                     # Log directory
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile                # Docker build file
├── Makefile                  # Makefile commands
├── go.mod                    # Go module definitions
├── go.sum                    # Go dependencies
└── README.md                 # This file
```

