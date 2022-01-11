FROM golang:1.17-alpine3.15 AS builder
WORKDIR /app
COPY  . .
RUN apk add curl
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY --from=builder /app/migrate ./migrate
COPY db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]