FROM alpine:3

ADD files /usr/bin/files

ENTRYPOINT ["/usr/bin/files"]
