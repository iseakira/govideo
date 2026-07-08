FROM golang:1.26-alpine AS builder

WORKDIR /go/src/ms-stream-api

COPY . .


ENV GO111MODULE=on


RUN go mod download
RUN go build -o /out/main main.go

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /app


COPY --from=builder /out/main ./main
COPY config ./config
COPY --from=builder /go/src/ms-stream-api/assets ./assets

RUN mkdir -p logs assets/media

ENTRYPOINT ["./main"]