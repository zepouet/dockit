FROM golang:1

MAINTAINER Nicolas MULLER <n.muller@treeptik.fr>

COPY /dockit-linux /dockit-linux

ENTRYPOINT ["/dockit-linux"]

