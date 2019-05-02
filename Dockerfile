FROM golang:1.12-alpine
MAINTAINER "Chavez <matthew@el-chavez.me>"

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
	&& wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-arm64.asc" \
	&& export GNUPGHOME="$(mktemp -d)" \
	&& gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
	&& gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
	&& rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
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
