FROM golang:1.15

ENV GO111MODULE=off

RUN go get -v github.com/markbates/refresh
RUN go get -u github.com/cosmtrek/air

ENV GO111MODULE=on

WORKDIR /go/src/github.com/reyhanfahlevi/poc-test