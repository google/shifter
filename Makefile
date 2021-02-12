.PHONY: test

all: test

test: clean
	go run . template -i ./_test/os-nginx-template.yaml -o ./out -k helm

clean: test
	rm -rf out

build:
	env GOOS=linux go build -v shifter
