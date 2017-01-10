FROM golang:1.7
MAINTAINER "Chavez <matthew@el-chavez.me>"

ENV APP_NAME=github.com/mtchavez/circlecli
ENV APP_DIR=$GOPATH/src/$APP_NAME
ENV APP_BIN=circlecli

ENV GOSU_VERSION=1.7
RUN set -x \
	&& wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$(dpkg --print-architecture)" \
	&& wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$(dpkg --print-architecture).asc" \
	&& export GNUPGHOME="$(mktemp -d)" \
	&& gpg --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
	&& gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
	&& rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
	&& chmod +x /usr/local/bin/gosu \
    && gosu nobody true

RUN groupadd -r circleuser && useradd -r -g circleuser circleuser

WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go get github.com/tools/godep && \
    godep restore
RUN go install -v

COPY ./docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["circlecli"]
