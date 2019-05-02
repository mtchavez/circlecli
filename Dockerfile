FROM golang:1.12-alpine

ENV APP_NAME=github.com/mtchavez/circlecli
ENV APP_DIR=$GOPATH/src/$APP_NAME
ENV APP_BIN=circlecli

RUN apk add \
  bash \
	git \
	mercurial \
	gnupg

ENV GOSU_VERSION=1.11
RUN set -x \
	&& wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-arm64" \
	&& chmod +x /usr/local/bin/gosu \
    && gosu nobody true

RUN addgroup -S circleuser && adduser -S circleuser -G circleuser

WORKDIR $APP_DIR
COPY . $APP_DIR

ENV GO111MODULE on
RUN go install -v

COPY ./docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["circlecli"]
