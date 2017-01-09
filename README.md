# circlecli

[![Latest Version](http://img.shields.io/github/release/mtchavez/circlecli.svg?style=flat-square)](https://github.com/mtchavez/circlecli/releases)
[![Build Status](https://travis-ci.org/mtchavez/circlecli.svg?branch=master)](https://travis-ci.org/mtchavez/circlecli)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/circlecli)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/circlecli)](https://goreportcard.com/report/github.com/mtchavez/circlecli)
[![Go Cover](http://gocover.io/_badge/github.com/mtchavez/circlecli)](http://gocover.io/github.com/mtchavez/circlecli)

## Docker

Installing via docker

```
docker build -t circlecli .
```

Using image to run CLI commands, can pass in `env-file` or `env` flags to be used
by the CLI.

```
docker run -it --env CIRCLECI_TOKEN=<TOKEN> --rm circlecli circlecli
```
