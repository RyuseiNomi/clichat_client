FROM golang:1.14

RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN go get github.com/gorilla/websocket