FROM golang:alpine
MAINTAINER Alexey Vakulich "soulSpringg@gmail.com"

# no need to redefine $GOPATH
ENV SRC_DIR=/go/src/github.com/interviewr/interviewr-server

ADD . $SRC_DIR
WORKDIR $SRC_DIR

RUN apk add --update --no-cache git
RUN mkdir -p /bin
RUN go get ./cmd/interviewr-server
RUN go build -o ./bin/entry ./cmd/interviewr-server

CMD ["./bin/entry"]

EXPOSE 8090
