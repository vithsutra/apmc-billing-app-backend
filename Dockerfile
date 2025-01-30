FROM golang:latest AS build-stage

WORKDIR /app

COPY . .

RUN go build -o ./bin/server ./cmd/*

FROM debian:stable

WORKDIR /app

RUN mkdir font-family

COPY --from=build-stage /app/bin/server .
COPY --from=build-stage /app/font-family ./font-family

CMD ["./server"]
