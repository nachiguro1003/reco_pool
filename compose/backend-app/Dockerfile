FROM golang:1.12.6-alpine3.9
LABEL maintainer=yoseiito
WORKDIR /go/src/github.com/nachiguro1003/reco-pool/
ENV GO111MODULE=on
RUN apk add --no-cache \
        git \
        && go get github.com/pilu/fresh
CMD ["/bin/true"]
