FROM golang:latest AS build-stage

WORKDIR /app

COPY . .

RUN go build -o ./bin/server ./cmd/*

FROM debian:stable

WORKDIR /app

COPY --from=build-stage /app/bin/server .

CMD ["./server"]
