FROM golang:1.11

EXPOSE 8080

WORKDIR /go/src/github.com/pchico83/observability-tests/tasks/worker
ADD . .

RUN go build -o /usr/local/bin/worker

CMD /usr/local/bin/worker

