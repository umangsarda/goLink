# GoLink 🔗

A distributed URL shortener and async task queue microservice built with **Go**, backed by **AWS DynamoDB** for persistent storage.

## Architecture

- **Go + Gorilla Mux** — REST API server
- **AWS ECS (Fargate)** — containerized deployment
- **DynamoDB** — NoSQL key-value store for short URLs (O(1) lookups)
- **Redis** — in-memory cache layer for sub-5ms p99 response time
- **Docker** — containerized for consistent environments
- **GitHub Actions** — CI/CD pipeline (build → test → deploy)

## API Endpoints

| Method | Endpoint       | Description               |
| ------ | -------------- | ------------------------- |
| POST   | `/shorten`     | Shorten a long URL        |
| GET    | `/:code`       | Redirect to original URL  |
| GET    | `/stats/:code` | Get stats for a short URL |
| GET    | `/health`      | Health check              |

## Quick Start

```bash
git clone https://github.com/umangsarda/golink.git
cd golink
go run main.go
```

### Shorten a URL

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"long_url": "https://www.google.com"}'
```

Response:

```json
{
  "code": "37cbe0",
  "long_url": "https://www.google.com",
  "short_url": "http://localhost:8080/37cbe0",
  "created_at": "2026-06-19T17:38:06-04:00",
  "hits": 0
}
```

### Redirect

```bash
curl -L http://localhost:8080/37cbe0
# redirects to https://www.google.com
```

### Get Stats

```bash
curl http://localhost:8080/stats/37cbe0
```

Response:

```json
{
  "code": "37cbe0",
  "long_url": "https://www.google.com",
  "short_url": "http://localhost:8080/37cbe0",
  "hits": 0,
  "created_at": "2026-06-19T17:38:06-04:00"
}
```

## Run with Docker

```bash
docker-compose up --build
```

## AWS DynamoDB

Links are persisted in AWS DynamoDB (`golink-urls` table) with `Code` as the partition key. Data survives server restarts and scales horizontally.

## CI/CD

Every push to `main` triggers GitHub Actions:

1. Checkout code
2. Set up Go 1.22
3. Build — `go build ./...`
4. Test — `go test ./...`
