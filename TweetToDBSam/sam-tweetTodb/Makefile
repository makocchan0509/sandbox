.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	rm -rf ./tweetToDynamo/tweetToDynamo
	rm -rf ./dynamoToS3/dynamoToS3

build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
	GOOS=linux GOARCH=amd64 go build -o tweetToDynamo/tweetToDynamo ./tweetToDynamo
	GOOS=linux GOARCH=amd64 go build -o dynamoToS3/dynamoToS3 ./dynamoToS3

package:
	sam package --template-file ./template.yaml --output-template-file ./output-template.yaml --s3-bucket masem.storage.sam.deploy

deploy:
	sam deploy --template-file ./output-template.yaml --stack-name go-lambda --capabilities CAPABILITY_IAM

