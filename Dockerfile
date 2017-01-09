FROM golang:1.7
MAINTAINER "Chavez <matthew@el-chavez.me>"

ENV APP_NAME=github.com/mtchavez/circlecli
ENV APP_DIR=$GOPATH/src/$APP_NAME
ENV APP_BIN=circlecli

WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go get github.com/tools/godep && \
    godep restore
RUN go install -v

CMD ["circlecli"]
