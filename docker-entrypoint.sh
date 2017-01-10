#!/bin/bash
set -e

# first arg is `-f` or `--some-option`
if [ "${1#-}" != "$1" ] || [ ! `command -v $1` ]; then
	set -- circlecli "$@"
fi

# allow the container to be started with `--user`
if [ "$1" = 'circlecli' -a "$(id -u)" = '0' ]; then
	chown -R circleuser .
	exec gosu circleuser "$0" "$@"
fi

exec "$@"
