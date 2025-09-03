
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download

# Generate openapi defitions
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
RUN swag init -g ./cmd/main.go

RUN go build -o api ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api .
CMD ["./api"]