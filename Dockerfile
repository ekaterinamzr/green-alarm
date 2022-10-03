# syntax=docker/dockerfile:1

# Build

FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

# Deploy 

FROM scratch

COPY --from=builder /app/config/config.yml /config/config.yml
COPY --from=builder /app/main /main

CMD ["/main"]