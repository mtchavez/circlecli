FROM golang:1.12-alpine as build

ENV APP_NAME=github.com/mtchavez/circlecli
ENV APP_DIR=$GOPATH/src/$APP_NAME
ENV APP_BIN=circlecli

RUN apk add \
	git \
	mercurial

WORKDIR $APP_DIR
COPY . $APP_DIR

ENV GO111MODULE on
RUN go install -v


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=build /go/bin/circlecli /usr/local/bin

ENTRYPOINT ["circlecli"]
CMD ["rb"]
