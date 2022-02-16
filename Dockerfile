# syntax=docker/dockerfile:1

##
## Build Shifter Server
##
FROM golang:1.17.7 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . ./

RUN go build -o /shifter

##
## Deploy Shifter Server
##
FROM gcr.io/distroless/base-debian10

ARG serverPort=8080
ENV env_serverPort=$serverPort


WORKDIR /

COPY --from=build /shifter /shifter

EXPOSE $env_serverPort
#$serverPort

USER nonroot:nonroot

CMD ./shifter server 