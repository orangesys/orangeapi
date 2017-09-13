FROM golang:1.8.3 as compile-stage
WORKDIR /go/src/github.com/orangesys/orangeapi
COPY . .
RUN make build

FROM alpine:3.6
MAINTAINER gavin zhou <gavin.zhou@gmail.com>
COPY --from=compile-stage dist/orangeapi_linux-amd64 /

RUN echo "==> Installing ..." \
  && apk add --no-cache ca-certificates

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
