# build stage
FROM node:lts-alpine as build-stage

# 
ENV ENV_JSON_FILE=./src/env-config.json

# Configure JQ for env-config.jon file Parsing
ENV JQ_VERSION=1.6
RUN wget --no-check-certificate https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 -O /tmp/jq-linux64
RUN cp /tmp/jq-linux64 /usr/bin/jq
RUN chmod +x /usr/bin/jq

WORKDIR /usr/src/app/
COPY package*.json ./
COPY . .

# ReParse the env-config.json file.
RUN jq 'to_entries | map_values({ (.key) : ("$" + .key) }) | reduce .[] as $item ({}; . + $item)' $ENV_JSON_FILE > ./src/env-config.tmp.json && mv ./src/env-config.tmp.json $ENV_JSON_FILE

# Build Shifter UI
RUN npm install
RUN npm run build

# Create Fixed Backup of Assets Folder
RUN ls 
RUN pwd
COPY /dist/assets /dist/assets_originals


# production stage
# nginx state for serving content
FROM nginx:alpine as production-stage

RUN apk add --update coreutils

# Add a user how will have the rights to change the files in code
RUN addgroup -g 1500 nginxusers 
RUN adduser --disabled-password -u 1501 nginxuser nginxusers 

# Configure ngnix server
COPY nginx/default.conf /etc/nginx/nginx.conf
WORKDIR /code
COPY --from=build-stage /usr/src/app/dist .

# Configure web-app for environment variable usage
WORKDIR /

# VueJS JS File Location
ENV JSFOLDER=/code/assets/*.js

COPY docker_entrypoint.sh .
#COPY generate_env-config.sh .

RUN chown nginxuser:nginxusers docker_entrypoint.sh
#RUN chown nginxuser:nginxusers generate_env-config.sh
RUN chmod 777 docker_entrypoint.sh 
#RUN chmod 777 generate_env-config.sh
RUN chown -R nginxuser:nginxusers /code
RUN chown -R nginxuser:nginxusers /etc/nginx
RUN chown -R nginxuser:nginxusers /tmp
RUN chmod 777 /code
RUN chmod 777 /tmp
RUN chmod 777 /etc/nginx

USER nginxuser

EXPOSE 8080:8080
CMD ["/bin/sh","docker_entrypoint.sh"]