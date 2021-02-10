.PHONY: test

all: test build

test: clean
	go run . template -i ./_test/os-nginx-template.yaml -o ./out -k helm

clean: test
	rm -rf out

build:
	go build

