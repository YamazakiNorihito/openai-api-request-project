# syntax=docker/dockerfile:1

FROM golang:1.20.4-alpine  AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]