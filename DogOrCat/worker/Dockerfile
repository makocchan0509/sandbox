# Worker app docker.
FROM golang:rc-buster

RUN mkdir -p /go/src/projects/DogOrCat/worker
WORKDIR /go/src/projects/DogOrCat/worker
COPY . .

RUN go get github.com/go-sql-driver/mysql && go get github.com/streadway/amqp &&  GOOS=linux GOARCH=amd64 go build -o worker main.go

CMD ["./worker"]
