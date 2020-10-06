FROM golang:1.15.2-alpine3.12

RUN apk add --no-cache go make

ADD . /app
WORKDIR /app

RUN make

ENTRYPOINT ["/app/files"]
