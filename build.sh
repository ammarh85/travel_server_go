#!/usr/bin/env bash

# Fail immediately if any of the commands below fail.
set -x

# Display the existing Go info before we change it.
go version
echo GOROOT=$GOROOT
echo GOPATH=$GOPATH

# Install a newer version of Go.
#GOGZ=go1.6.2.linux-amd64.tar.gz
#curl -O "https://storage.googleapis.com/golang/$GOGZ"
#tar -zxf $GOGZ
#rm $GOGZ

#export GOROOT=/usr/local/go
#export GOPATH=$GOROOT/bin
#rm -rf $GOROOT
#mv go $GOROOT
#export PATH=$GOPATH:$PATH

# Make sure we are now using the newer version.
go version
echo GOROOT=$GOROOT
echo GOPATH=$GOPATH
echo "Hello World Path!"
# Move the application into the correct package for $GOPATH.
#DEST=$GOROOT/src/github.com/your/package
#rm -rf $DEST
#mkdir -p $DEST
#cp -r * $DEST   # NOTE: must be 'cp', not 'mv' here.

# Install/update dependencies.
#go get -u gopkg.in/redis.v3
#go get -u github.com/Sirupsen/logrus
#go get -u gopkg.in/yaml.v2
go get -u github.com/gorilla/mux
echo "Hello World Post MUX!"

# Build the server. Beanstalk will look for an 'application' executable.
#cd $DEST
#go build -o application
#chmod +x application
#mv application /usr/bin