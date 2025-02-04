FROM golang:1.23.5 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

ARG APP_PORT=8080

ENV APP_PORT=${APP_PORT}
EXPOSE ${APP_PORT}
ENTRYPOINT ["./main"]
