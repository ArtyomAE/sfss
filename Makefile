all: build

build:
	go build .

install: build
	cp -f git-fit /usr/local/bin/git-fit

deps:
	go get -u github.com/mitchellh/goamz/aws
	go get -u github.com/mitchellh/goamz/s3
	go get -u github.com/cheggaaa/pb
	go get -u github.com/github.com/smartystreets/goconvey/convey

unittests:
	go test github.com/dailymuse/git-fit/transport github.com/dailymuse/git-fit/util

integrationtests: build
	./integration.py; rm -rf integration

test: unittests integrationtests
