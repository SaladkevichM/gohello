SHELL := /bin/bash

# The name of the executable
TARGET := 'stack_bin'

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

default: 
	go build -o $(TARGET) -v ./src/main/stack/stack.go 

#install:
#	go install ./src/main/stack/stack.go 

clean:
	go clean
	rm -f ./$(TARGET)
 
run: clean 
	go build -o $(TARGET) ./src/main/stack/stack.go 
	./$(TARGET)

#test: 
#    go test -v ./src/main/stack/test_stack.go

#docker-build:
#    docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

#deps:
#    go get github.com/markbates/goth
#    go get github.com/markbates/pop
