FROM golang:1.11

RUN go get github.com/vapor-ware/ksync/cmd/ksync

WORKDIR /go/src/github.com/pchico83/ksync
ADD . .

RUN go build -o /usr/local/bin/ksync-server

CMD /usr/local/bin/ksync-server

