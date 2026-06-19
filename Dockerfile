# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o golink .

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/golink .
EXPOSE 8080
CMD ["./golink"]