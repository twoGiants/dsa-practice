FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o dsa .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dsa .

CMD ["./dsa"]