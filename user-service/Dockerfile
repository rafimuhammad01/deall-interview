FROM golang:1.19.3-alpine3.16
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o user-service

ENTRYPOINT ["/app/user-service"]

