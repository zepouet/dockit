FROM golang

MAINTAINER Nicolas MULLER <n.muller@treeptik.fr>

COPY /dockit-linux /dockit-linux

ENTRYPOINT ["/dockit-linux"]

