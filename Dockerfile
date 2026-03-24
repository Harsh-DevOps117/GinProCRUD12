FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY .env .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

FROM alpine:3.21

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .

RUN chown appuser:appgroup /app/main

USER appuser

EXPOSE 8000

CMD ["./main"]
