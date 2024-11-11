
FROM golang:1.20 as builder

WORKDIR /app


COPY . .


RUN go mod init hola-mundo && go mod tidy


FROM golang:1.20


WORKDIR /app


COPY --from=builder /app /app


EXPOSE 8080


CMD ["go", "run", "main.go"]

