#!/usr/bin/env bash

# Fail immediately if any of the commands below fail.
set -x

# Display the existing Go info
go version
echo GOROOT=$GOROOT
echo GOPATH=$GOPATH

# Install/update dependencies.
go get -u github.com/gorilla/mux
go get -u github.com/aws/aws-sdk-go/aws

# Build the server. Beanstalk will look for an 'application' executable.
go build -o bin/application application.go