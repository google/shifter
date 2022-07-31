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


#! /bin/bash

# Automate the creation of OKD cluster 
# Scenario 1 : Script is running on a VM with the relevent permissions attached 
set -e 
source $(dirname "$0")/variables.sh

# Generate SSH key to attach to the bastion host and enable passwordless access 
function generate_ssh() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Generate SSH function starts ------------" >> ${LOG_FILE}
if [ -f $SSH_PATH ] 
then
  echo "$(date +'%Y-%m-%d %H:%M:%S'): SSH Key is already present." >> ${LOG_FILE}
else  
echo "$(date +'%Y-%m-%d %H:%M:%S'): Generating SSH keys $SSH_PATH " >> ${LOG_FILE}
ssh-keygen -b 2048 -t rsa -f $SSH_PATH -q -N "" <<< $'\ny' >/dev/null 2>&1
fi
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Generate SSH function Ends ------------" >> ${LOG_FILE} 
}

# function to provision infrastructure via terraform 
function provision_infra() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision Infra function starts ------------" >> ${LOG_FILE}
cd ../terraform
echo "$(date +'%Y-%m-%d %H:%M:%S'): Running terraform to provision the backend infrastructure" >> ${LOG_FILE}
terraform init  >> ${LOG_FILE}
terraform plan -out plan.out -var="gce_ssh_user=$SSH_USER" -var="gce_ssh_pub_key_file=$SSH_PUB_FILE" -var="region=$REGION" -var="ssh_user=$SSH_USER" >> ${LOG_FILE}
terraform apply plan.out >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision Infra function Ends ------------" >> ${LOG_FILE}
}

# function to copy the host and the ssh files
function copy_files() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Copy files function starts ------------" >> ${LOG_FILE}
BASTION_HOST=$(terraform output bastion | tr -d '"')
PROJECT=$(terraform output project_id | tr -d '"')
MASTER=$(terraform output master | tr -d '"')
LB_IP=$(terraform output google_compute_address | tr -d '"')
echo "$(date +'%Y-%m-%d %H:%M:%S'): Copying ssh key and host file" >> ${LOG_FILE}
gcloud compute scp  --project=$PROJECT --zone=$ZONE $SSH_PATH $SSH_USER@$BASTION_HOST:~/.ssh/id_rsa >> ${LOG_FILE}
gcloud compute scp  --project=$PROJECT --zone=$ZONE ./inventory/ansible-hosts $SSH_USER@$BASTION_HOST:/home/$SSH_USER >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'): Successfully copied ssh key and host file" >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Copy files function ends ------------" >> ${LOG_FILE}
}

# function to provision the  OKD cluster
function provision_cluster() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision cluster function starts ------------" >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):SSH into the bastion host and run ansible scripts" >> ${LOG_FILE}
gcloud compute ssh --project=$PROJECT --zone=$ZONE $SSH_USER@$BASTION_HOST >> ${LOG_FILE} << EOF 
function run_ansible() {
if [ -d openshift-ansible ] 
then
  echo "$(date +'%Y-%m-%d %H:%M:%S'): Fork already cloned."
else
 echo "$(date +'%Y-%m-%d %H:%M:%S'): Clone Openshift-ansible Github Repository" 
 git clone https://github.com/openshift/openshift-ansible.git
fi
mv /home/$SSH_USER/ansible-hosts /home/$SSH_USER/openshift-ansible/inventory
cd openshift-ansible
git checkout $OKD_VERSION
ansible-playbook -i inventory/ansible-hosts playbooks/prerequisites.yml 
ansible-playbook -i inventory/ansible-hosts playbooks/deploy_cluster.yml
echo "$(date +'%Y-%m-%d %H:%M:%S'): Exiting from the bastion host."
}
run_ansible
EOF
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision cluster function ends ------------" >> ${LOG_FILE}
}


## function to deploy manifest on the cluster 
function deploy_manifest() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- deploy manifest function starts ------------" >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):SSH into the master host to deploy manifest files" >> ${LOG_FILE}

gcloud compute ssh --project=$PROJECT --zone=$ZONE $SSH_USER@$MASTER >> ${LOG_FILE} << EOF
function deploy_boa() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Deploy BOA function starts ------------"

if [ -d bank-of-anthos ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S'): Fork already cloned."
else
 echo "$(date +'%Y-%m-%d %H:%M:%S'): Clone bank-of-anthos Github Repository" 
 git clone https://github.com/avinashkumar1289/bank-of-anthos.git
fi
oc adm policy add-scc-to-user privileged system:serviceaccount:default:default
oc apply -f bank-of-anthos/kubernetes-manifest/jwt/jwt-secret.yaml
oc apply -f bank-of-anthos/kubernetes-manifest/.
echo "############################################################"
echo "Waiting for  60 seconds for workloads to be ready..."
echo "############################################################"
sleep 60s
oc get pods
echo "############################################################"
echo "Endpoint"
sleep 60s
echo "############################################################"
oc get service frontend | awk '{print $4}'
echo "$(date +'%Y-%m-%d %H:%M:%S'): Exiting from the master node."
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Deploy BOA function Ends ------------"
} 
deploy_boa
EOF
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- deploy manifest function ends ------------" >> ${LOG_FILE}
}

## function to set configuraiton for the OKD cluster 
function okd_config() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):--------okd set configurations starts ------------" >> ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):SSH into the master host to set configurations" >> ${LOG_FILE}

gcloud compute ssh --project=$PROJECT --zone=$ZONE $SSH_USER@$MASTER >> ${LOG_FILE} << EOF
function set_config() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Set Configfunction starts ------------"
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Assign cluster role to the user shifter ------------"
oc adm policy add-cluster-role-to-user cluster-admin shifter
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Generating the Bearer token ------------"
curl -u shifter:shifter -kv '$LB_IP:8443/oauth/authorize?client_id=openshift-challenging-client&response_type=token' -skv -H "X-CSRF-Token: xxx"
echo "$(date +'%Y-%m-%d %H:%M:%S'): Exiting from the master node."
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Set configuration function Ends ------------"
} 
set_config
EOF
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- okd set configurations function ends ------------" >> ${LOG_FILE}
}

mkdir -p ${SSH_PATH} || true
mkdir -p ${LOG_PATH}
touch ${LOG_FILE}

generate_ssh
provision_infra
copy_files
provision_cluster
deploy_manifest
okd_config