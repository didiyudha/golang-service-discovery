FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get github.com/hashicorp/consul/api

ENV SOURCE=/go/src/github.com/didiyudha/consul-go/
COPY . ${SOURCE}

RUN cd ${SOURCE}product-service/ && CGO_ENABLED=0 go build

WORKDIR ${SOURCE}
CMD ${SOURCE}product-service/product-service
EXPOSE 8100