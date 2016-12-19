GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint \
		github.com/tools/godep
default: test

ci: deps test

deps:
	@go get -u -v $(GOPKGS)
	@if [ `which godep` ] && [ -f ./Godeps/Godeps.json ]; then godep restore; fi

build: pre_test
	@./script/build

pre_test:
	@./script/pre_test

test: pre_test
	@./script/test

PHONY: ci deps build pre_test test
