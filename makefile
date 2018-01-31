SHELL := /bin/bash

# The name of the executable
TARGET := $(shell echo $${PWD\#\#*/})

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

default: 
	echo TARGET
	go build ./src/main/stack/stack.go

install:
	go install

clean:
	go clean	