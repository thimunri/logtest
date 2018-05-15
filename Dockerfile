FROM golang:1.9.1

RUN mkdir -p /go/src/github.com/thimunri/logtest
RUN mkdir -p /go/bin/
RUN mkdir -p /go/pkg

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/thimunri/logtest
ADD . /go/src/github.com/thimunri/logtest
RUN glide install
RUN go build -o /app/main
