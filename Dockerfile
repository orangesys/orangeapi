FROM golang:1.8.3 as compile-stage 
WORKDIR /go/src/github.com/orangesys/orangeapi
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  && GOARCH=amd64 GOOS=linux CGO_ENABLED=0 \
  go build -ldflags '-s -w' -o dist/orangeapi_linux-amd64 .

FROM alpine:3.6
MAINTAINER gavin zhou <gavin.zhou@gmail.com>
COPY --from=compile-stage dist/orangeapi_linux-amd64 /

RUN echo "==> Installing ..." \
  && apk add --no-cache ca-certificates

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
