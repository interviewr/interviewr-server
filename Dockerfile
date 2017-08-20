FROM golang:alpine
MAINTAINER Alexey Vakulich "soulSpringg@gmail.com"

# no need to redefine $GOPATH
ENV SRC_DIR=/go/src/interviewr/interviewr-server

ADD . $SRC_DIR
WORKDIR $SRC_DIR

RUN apk add --update --no-cache git
RUN mkdir -p /bin
RUN go get ./src/cmd
RUN go build -o ./bin/entry ./src/cmd

CMD ["./bin/entry"]

EXPOSE 8090
