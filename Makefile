GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint \

default: test

ci: deps test

deps:
	@go get -v $(GOPKGS)

build: pre_test
	@./script/build

pre_test:
	@./script/pre_test

test: pre_test
	@./script/test

PHONY: ci deps build pre_test test
