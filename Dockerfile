FROM golang:1.7.3-alpine
MAINTAINER Dany Marcoux <danymarcoux@gmail.com>

RUN mkdir -p "$GOPATH/src/github.com/dmarcoux/bureaucrat"
WORKDIR $GOPATH/src/github.com/dmarcoux/bureaucrat
ADD / .

RUN go install

ENTRYPOINT $GOPATH/bin/bureaucrat
