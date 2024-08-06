FROM golang:1.22.1

RUN go install github.com/golang/dep/cmd/dep@latest

WORKDIR /go/src/github.com/abhishekkushwahaa/golang-microservice
COPY . .

RUN dep ensure
RUN go get github.com/cespare/reflex