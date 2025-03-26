FROM golang:latest AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/main /app/bin/main

COPY .env .env

RUN chmod +x /app/bin/main

EXPOSE 8080

CMD ["/app/bin/main"]
