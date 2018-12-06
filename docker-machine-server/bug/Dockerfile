FROM golang:1.11

RUN curl -L https://github.com/docker/machine/releases/download/v0.15.0/docker-machine-$(uname -s)-$(uname -m) > /usr/local/bin/docker-machine
RUN chmod +x /usr/local/bin/docker-machine

WORKDIR /go/src/github.com/pchico83/observability-tests/docker-machine-server/bug
ADD . .

RUN go build -o /usr/local/bin/docker-machine-server

ENV AWS_REGION us-west-1

CMD /usr/local/bin/docker-machine-server

