FROM alpine:3

RUN mkdir /app
ADD public /app/public
ADD files /usr/bin/files

WORKDIR /app

ENTRYPOINT ["/usr/bin/files"]
