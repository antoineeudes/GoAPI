FROM golang:1.10

RUN mkdir -p /go/src/API/
WORKDIR /go/src/API/
COPY . /go/src/API/

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/handlers
RUN go get github.com/jinzhu/gorm
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt