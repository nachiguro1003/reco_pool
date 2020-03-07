FROM golang:1.12.6
LABEL maintainer=yoseiito
WORKDIR /go/src/github.com/reco-pool
ENV GO111MODULE=on
RUN apk add --no-cache \
        git \
    && go get github.com/pilu/fresh
EXPOSE 8080
CMD ["fresh"]
