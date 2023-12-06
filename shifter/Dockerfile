# syntax=docker/dockerfile:1
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


##
## Build Shifter Server
##
FROM golang:1.21.5 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . ./

RUN go build -o /shifter

##
## Deploy Shifter Server
##
FROM gcr.io/distroless/base-debian11 

ARG serverPort=8080
ENV env_serverPort=$serverPort

WORKDIR /shifter
COPY --from=build /shifter ./
EXPOSE 8080
CMD ["/shifter"]
