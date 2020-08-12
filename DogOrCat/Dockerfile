# Web app docker.
FROM golang:rc-buster

RUN mkdir -p /go/src/projects/DogOrCat
WORKDIR /go/src/projects/DogOrCat
COPY web web
WORKDIR /go/src/projects/DogOrCat/web

RUN go get github.com/go-sql-driver/mysql && go get github.com/streadway/amqp &&  GOOS=linux GOARCH=amd64 go build -o webapp main.go

EXPOSE 8080

CMD ["./webapp"]
