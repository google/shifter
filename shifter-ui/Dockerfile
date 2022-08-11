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

# build stage
FROM node:lts-alpine as build-stage
ENV ENV_JSON_FILE=./src/env-config.json
ENV JQ_VERSION=1.6

# Configure JQ for env-config.jon file Parsing
RUN wget --no-check-certificate https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 -O /tmp/jq-linux64 && \
    cp /tmp/jq-linux64 /usr/bin/jq && \
    chmod +x /usr/bin/jq

WORKDIR /usr/src/app/
#COPY package*.json ./
COPY . .

# ReParse the env-config.json file.
RUN jq 'to_entries | map_values({ (.key) : ("$" + .key) }) | reduce .[] as $item ({}; . + $item)' $ENV_JSON_FILE > ./src/env-config.tmp.json && mv ./src/env-config.tmp.json $ENV_JSON_FILE

# Build Shifter UI
RUN npm install && \
    npm run build

# production stage
# nginx state for serving content
FROM nginx:alpine as production-stage

RUN apk add --update coreutils

# Add a user how will have the rights to change the files in code
RUN addgroup -g 1500 nginxusers && \
    adduser --disabled-password -u 1501 nginxuser nginxusers 

# Configure ngnix server
WORKDIR /code
COPY nginx/default.conf /etc/nginx/nginx.conf
COPY --from=build-stage /usr/src/app/dist .
COPY --from=build-stage /usr/src/app/dist/assets ./assets_original

# Configure web-app for environment variable usage
WORKDIR /

# VueJS JS File Location
ENV JSFOLDER=/code/assets/*.js

COPY docker_entrypoint.sh .
#COPY generate_env-config.sh .

RUN chown nginxuser:nginxusers docker_entrypoint.sh && \
    chmod +x docker_entrypoint.sh && \
    chown -R nginxuser:nginxusers /code && \
    chown -R nginxuser:nginxusers /etc/nginx && \
    #chown -R nginxuser:nginxusers /tmp && \
    chown nginxuser:nginxusers /code 

USER nginxuser

EXPOSE 8080:8080
CMD ["/bin/sh","docker_entrypoint.sh"]
