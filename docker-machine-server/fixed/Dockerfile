FROM golang:1.11

RUN echo install syncthing

WORKDIR /go/src/github.com/pchico83/docker-machine-server
ADD . .

RUN go build -o /usr/local/bin/docker-machine-server

CMD /usr/local/bin/docker-machine-server

