FROM alpine:3.4

ENV GOPATH /go
COPY . /go/src/github.com/koudaiii/sample-go-api

RUN apk add --no-cache --update --virtual=build-deps g++ git go mercurial make curl \
    && cd /go/src/github.com/koudaiii/sample-go-api \
    && make \
    && cp bin/sample-go-api /server \
    && cd / \
    && rm -rf /go \
    && apk del build-deps

EXPOSE 8080

CMD ["/server"]
