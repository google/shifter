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
FROM debian:11-slim

ARG serverPort=8080
ENV env_serverPort=$serverPort


WORKDIR /shifter
COPY --from=build /shifter ./
EXPOSE 8080
CMD ["./shifter", "server"]