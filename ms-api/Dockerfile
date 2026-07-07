FROM golang:1.26-alpine AS builder
WORKDIR /go/src/ms-api
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /out/main .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates ffmpeg
WORKDIR /app

COPY --from=builder /out/main ./main
COPY config ./config
COPY scripts ./scripts
COPY assets ./assets

RUN mkdir -p logs
ENTRYPOINT ["./main"]