FROM debian:buster-slim

RUN apt-get update && apt-get install -y wget

RUN wget https://get.helm.sh/helm-v3.5.2-linux-amd64.tar.gz && tar -xvzf helm-v3.5.2-linux-amd64.tar.gz && mv linux-amd64/helm /usr/local/bin && rm -rf linux-amd64
