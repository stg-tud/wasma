FROM golang:1.17-alpine

WORKDIR /usr/wasma

COPY . /usr/wasma

RUN apk add --no-cache nano
RUN apk add --no-cache make
RUN make build

ENV PATH="/usr/wasma/bin:${PATH}"
ENV PATH="/usr/wasma/tools/scripts:${PATH}"

WORKDIR /usr/wasma/bin