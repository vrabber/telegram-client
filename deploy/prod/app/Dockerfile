FROM golang:1.23 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o telegram-client ./cmd/client/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/telegram-client .
CMD [ "./telegram-client" ]