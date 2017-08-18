FROM golang:alpine
MAINTAINER Alexey Vakulich "soulSpringg@gmail.com"

ENV SRC_DIR=/go/src/interviewr

ADD . $SRC_DIR
WORKDIR $SRC_DIR

RUN apk add --update --no-cache git
RUN mkdir -p /bin
RUN go get ./src/cmd
RUN go build -o ./bin/entry ./src/cmd

CMD ["./bin/entry"]

EXPOSE 8090
