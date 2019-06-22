FROM golang:1.10

RUN mkdir -p /go/src/API/
WORKDIR /go/src/API/
COPY . /go/src/API/

RUN go get -v
