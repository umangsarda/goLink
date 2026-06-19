# GoLink 🔗

A distributed URL shortener and async task queue microservice built with **Go**, deployed on **AWS ECS (Fargate)**.

## Architecture

- **Go + Gorilla Mux** — REST API server
- **AWS ECS (Fargate)** — containerized deployment
- **DynamoDB** — NoSQL key-value store for short URLs
- **Redis** — in-memory cache for sub-5ms p99 lookups
- **Docker** — containerized for consistent environments
- **GitHub Actions** — CI/CD pipeline (build → test → deploy)

## API Endpoints

| Method | Endpoint   | Description              |
| ------ | ---------- | ------------------------ |
| POST   | `/shorten` | Shorten a long URL       |
| GET    | `/:code`   | Redirect to original URL |
| GET    | `/health`  | Health check             |

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

### Use the short URL

```bash
curl -L http://localhost:8080/{code}
```

## CI/CD

Every push to `main` triggers GitHub Actions — builds, tests, and validates the Go binary automatically.
