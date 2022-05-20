FROM alpine:3

RUN mkdir /app
ADD public /app/public
ADD files-linux-amd64 /usr/bin/files

WORKDIR /app

ENTRYPOINT ["/usr/bin/files"]
