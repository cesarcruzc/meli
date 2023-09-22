FROM golang:alpine

COPY . /app
WORKDIR /app

EXPOSE 8888

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

CMD ["./main"]