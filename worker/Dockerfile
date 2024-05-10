# Use the official Golang image as a base image
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o worker .

CMD ["./worker"]

