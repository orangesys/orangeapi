FROM alpine:3.4
MAINTAINER gavin zhou <gavin.zhou@gmail.com>
ADD bin/orangeapi /orangeapi

RUN echo "==> Installing ..." \
  && apk add --no-cache ca-certificates

COPY docker-entrypoint.sh /docker-entrypoint.sh
EXPOSE 1323
ENTRYPOINT ["/docker-entrypoint.sh"]
