.PHONY: deps clean build

deps:
	go get

test:
	go test -run ''

clean: 
	rm -rf ../build/webapp

build:
	GOOS=linux GOARCH=amd64 go build -o ../build/webapp ./main.go

release:
	scp -i ../key/masem-ec2-key.pem ../build/webapp ec2-user@$(host):/home/ec2-user/go/src/projects/outputlog