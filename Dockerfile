FROM golang:1.24.6 AS builder
WORKDIR /app
COPY . .
WORKDIR /app/app/echo-server
RUN go build -o echo-server main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY .env .
COPY --from=builder /app/app/echo-server/echo-server /echo-server
CMD ["/echo-server"]