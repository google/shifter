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
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Generate SSH function starts ------------" 2>&1 | tee  ${LOG_FILE}
if [ -f $SSH_PATH ] 
then
  echo "$(date +'%Y-%m-%d %H:%M:%S'): SSH Key is already present." 2>&1 | tee  ${LOG_FILE}
else  
echo "$(date +'%Y-%m-%d %H:%M:%S'): Generating SSH keys $SSH_PATH " 2>&1 | tee ${LOG_FILE}
ssh-keygen -b 2048 -t rsa -f $SSH_PATH -q -N "" <<< $'\ny' >/dev/null 2>&1
fi
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Generate SSH function Ends ------------" 2>&1 | tee ${LOG_FILE} 
}

# function to provision infrastructure via terraform 
function provision_infra() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision Infra function starts ------------" 2>&1 | tee ${LOG_FILE}
cd ../terraform
echo "$(date +'%Y-%m-%d %H:%M:%S'): Running terraform to provision the backend infrastructure" 2>&1 | tee ${LOG_FILE}
terraform init  2>&1 | tee ${LOG_FILE}
terraform plan -out plan.out -var="gce_ssh_user=$SSH_USER" -var="gce_ssh_pub_key_file=$SSH_PUB_FILE" -var="region=$REGION" 2>&1 | tee ${LOG_FILE}
terraform apply plan.out 2>&1 | tee ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision Infra function Ends ------------" 2>&1 | tee ${LOG_FILE}
}

# function to copy the host and the ssh files
function copy_files() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Copy files function starts ------------" 2>&1 | tee ${LOG_FILE}
BASTION_HOST=$(terraform output bastion | tr -d '"')
PROJECT=$(terraform output project_id | tr -d '"')
MASTER=$(terraform output master | tr -d '"')
echo "$(date +'%Y-%m-%d %H:%M:%S'): Copying ssh key and host file" 2>&1 | tee ${LOG_FILE}
gcloud compute scp  --project=$PROJECT --zone=$ZONE $SSH_PATH $SSH_USER@$BASTION_HOST:~/.ssh/id_rsa 2>&1 | tee ${LOG_FILE}
gcloud compute scp  --project=$PROJECT --zone=$ZONE ./inventory/ansible-hosts $SSH_USER@$BASTION_HOST:/home/$SSH_USER 2>&1 | tee ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'): Successfully copied ssh key and host file" 2>&1 | tee ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Copy files function ends ------------" 2>&1 | tee ${LOG_FILE}
}

# function that runs the ansible playbook to deploy a OKD cluster
function provision_cluster() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision cluster function starts ------------" 2>&1 | tee ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):SSH into the bastion host and run ansible scripts" 2>&1 | tee ${LOG_FILE}
gcloud compute ssh --project=$PROJECT --zone=$ZONE $SSH_USER@$BASTION_HOST > ${LOG_FILE} << EOF
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
EOF 
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- Provision cluster function ends ------------" 2>&1 | tee ${LOG_FILE}
}

## function to deploy bank of anthos manifest on the cluster 
function deploy_manifest() {
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- deploy manifest function starts ------------" 2>&1 | tee ${LOG_FILE}
echo "$(date +'%Y-%m-%d %H:%M:%S'):SSH into the master host to deploy manifest files" 2>&1 | tee ${LOG_FILE}
gcloud compute ssh --project=$PROJECT --zone=$ZONE $SSH_USER@$MASTER > ${LOG_FILE} << EOF
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
EOF
echo "$(date +'%Y-%m-%d %H:%M:%S'):-------- deploy manifest function ends ------------" 2>&1 | tee ${LOG_FILE}
}


mkdir -p ${SSH_PATH} || true
mkdir -p ${LOG_PATH} || true
touch ${LOG_FILE}

generate_ssh
provision_infra
copy_files
provision_cluster
deploy_manifest