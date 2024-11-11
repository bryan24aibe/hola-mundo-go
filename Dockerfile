
FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod init hola-mundo && go mod tidy && go build -o app

FROM alpine:latest

COPY --from=builder /app/app /app/app

WORKDIR /app
CMD ["./app"]
