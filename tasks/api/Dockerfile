FROM golang:1.11

EXPOSE 8080

WORKDIR /go/src/github.com/pchico83/observability-tests/tasks/api
ADD . .

RUN go build -o /usr/local/bin/api

CMD /usr/local/bin/api

