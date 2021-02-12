.PHONY: test

all: test

test: clean templateTest

templateTest:
	go run . template -i ./_test/os-nginx-template.yaml -o ./out -k helm

clean: 
	rm -rf out

build:
	env GOOS=linux go build -v shifter
