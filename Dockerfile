FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY server.go tipe.pdf ./

RUN go build -o server server.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server /app/tipe.pdf ./

EXPOSE 8888

CMD ["./server"]
