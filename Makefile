.PHONY: test

all: test

test: clean templateTest lint apply

lint:
	helm lint ./out

templateTest:
	go run . template -i ./_test/os-nginx-template.yaml -o ./out -k helm

clean: 
	kubectl delete ns test
	kubectl create ns test
	rm -rf out

build:
	env GOOS=linux go build -v shifter

apply:
	helm install  ./out -n test --generate-name
