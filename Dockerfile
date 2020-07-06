FROM golang:latest

COPY ./ /go/src/goChess

RUN chmod +x /go/src/goChess/entrypoint.sh

ENTRYPOINT /go/src/goChess/entrypoint.sh