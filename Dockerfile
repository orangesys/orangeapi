FROM alpine:3.6
MAINTAINER gavin zhou <gavin.zhou@gmail.com>
COPY dist/orangeapi_linux-amd64 /
COPY linux-amd64/helm /usr/sbin/

RUN echo "==> Installing ..." \
  && apk add --no-cache ca-certificates

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
