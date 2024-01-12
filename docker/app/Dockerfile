FROM golang:1.21.5-alpine3.19
RUN mkdir sesgo
WORKDIR /sesgo

RUN apk update --no-cache

COPY . .

RUN source .env
RUN go build -o main cmd/cli/main.go

CMD ["./main"]