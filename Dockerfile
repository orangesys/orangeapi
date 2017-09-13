FROM golang:1.8.3-alpine3.6 AS compile-stage
WORKDIR /go/src/github.com/orangesys/orangeapi
COPY . .
RUN make build

FROM scratch
MAINTAINER gavin zhou <gavin.zhou@gmail.com>
COPY --from=compile-stage /go/src/github.com/orangesys/orangeapi/dist/orangeapi_linux-amd64 /
COPY --from=compile-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
