FROM ubuntu:latest
FROM golang as builder

RUN apt update && apt full-upgrade
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go get  github.com/labstack/echo
RUN go get	github.com/dgrijalva/jwt-go 
RUN go get	github.com/labstack/gommon 
RUN go get	github.com/lib/pq 
RUN go get	github.com/mattn/go-colorable 
RUN go get	github.com/mattn/go-isatty
RUN go get	github.com/stretchr/testify
RUN go get	github.com/valyala/bytebufferpool 
RUN go get	github.com/valyala/fasttemplate
RUN go get	golang.org/x/crypto 
RUN go get	golang.org/x/net
RUN go get	golang.org/x/sys 
RUN go get	golang.org/x/text 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
EXPOSE 5000
CMD [ "/go/src/app/main" ]