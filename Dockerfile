###
# Mainflux MongoDB Reader Dockerfile
###

FROM golang:alpine
MAINTAINER Mainflux

ENV MONGO_HOST mongo
ENV MONGO_PORT 27017

###
# Install
###
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mainflux/mainflux-mongodb-reader
RUN cd /go/src/github.com/mainflux/mainflux-mongodb-reader && go install

###
# Run main command with dockerize
###
CMD mainflux-mongodb-reader -m $MONGO_HOST
