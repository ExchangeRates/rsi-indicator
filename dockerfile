FROM golang:alpine3.15 as build

WORKDIR /app

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -v ./cmd/server

FROM alpine:3.15

WORKDIR /app

COPY /config ./config
COPY --from=build /app/server ./server

EXPOSE 8081

CMD ["./server"]

