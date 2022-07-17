# Copyright 2021 Google LLC
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


#!/bin/bash

SSH_PATH=${HOME}/gcp_keys/id_rsa
SSH_PUB_FILE=${SSH_PATH}.pub
SSH_USER=avinashjha
REGION=europe-west1
ZONE=europe-west1-b
OKD_VERSION=release-3.11
LOG_PATH="${HOME}/logs-$(date +'%Y-%m-%d-%H:%M:%S')"
LOG_FILE="${LOG_PATH}/okd-logs-$(date +'%Y-%m-%d-%H:%M:%S')"