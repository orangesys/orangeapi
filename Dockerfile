FROM golang:1.8.3 as compile-stage
WORKDIR /go/src/github.com/orangesys/orangeapi
COPY . .
RUN make build

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
LABEL maintainer "gavin zhou <gavin.zhou@gmail.com>""
COPY --from=compile-stage /go/src/github.com/orangesys/orangeapi/dist/orangeapi_linux-amd64 /
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
