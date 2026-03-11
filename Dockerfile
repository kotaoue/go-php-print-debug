# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY tools/search_print_debug/go.mod ./
COPY tools/search_print_debug/main.go ./

RUN go build -o search-print-debug .

# Runtime stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/search-print-debug .

ENTRYPOINT ["/app/search-print-debug"]
