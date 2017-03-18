FROM golang:1.7

MAINTAINER mochi8k

RUN apt-get update && apt-get install -y redis-server
RUN go get -u -t github.com/mochi8k/aiteru-server

ADD wait-for-it.sh /go/src/github.com/mochi8k/aiteru-server

WORKDIR /go/src/github.com/mochi8k/aiteru-server

CMD ["sh", "wait-for-it.sh", "go", "run", "server.go"]
