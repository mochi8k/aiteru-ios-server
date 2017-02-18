FROM golang:1.7
MAINTAINER mochi8k
RUN go get -u -t github.com/mochi8k/aiteru-server
WORKDIR /go/src/github.com/mochi8k/aiteru-server
CMD ["go", "run", "server.go"]
