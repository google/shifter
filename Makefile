.PHONY: test

all: test

test: templateTest lint

lint:
	helm lint ./out

templateTest:
	go run . convert -t template -i ./_test/os-nginx-template.yaml -o ./out -k helm

clean:

build:
	env GOOS=linux go build -v shifter

apply:
	helm install  ./out -n test --generate-name
