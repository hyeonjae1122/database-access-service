FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# 명시적으로 cross-compile
RUN GOOS=linux GOARCH=amd64 go build -o server

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/server .

CMD ["./server"]
