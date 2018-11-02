FROM golang:1.11

RUN curl -L https://github.com/docker/machine/releases/download/v0.15.0/docker-machine-$(uname -s)-$(uname -m) > /usr/local/bin/docker-machine
RUN chmod +x /usr/local/bin/docker-machine

WORKDIR /go/src/github.com/pchico83/docker-machine-server
ADD . .

RUN go build -o /usr/local/bin/docker-machine-server

CMD /usr/local/bin/docker-machine-server

