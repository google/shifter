
## Introduction

Automate the provisioning of OKD 3.11 cluster on GCP 

## Folder Structure

1. This repository help you create a openshift 3.11 cluster on GCP. This can also be extended to install lower versions of openshfit cluster.
2. `script` contains the files to automate the deployment of relevent infrastrcture required on GCP. It also run the ansible playbook to deploy the OKD cluster.  


## Requirements
<ol>
<li>The terraform script needs the GCP owner role in order to provision the infrastrcture. If the users are running this from a GCP VM, a service account needs to be attached to the VM with the proper roles. If running the script outside of GCP, service account keys should be used to authenticate</li>
<li> A valid billing account details are rquired before deploying the script. The billing details needs to be updated in the terraform tfvars files</li>
<li>A cloud storage bucket needs to be provisioned and updated in the backend.tfvars files </li>
<li>A valid domain is required to deploy the openshift cluster and access the console</li>
<li>scripts were tested in a linux machine.</li>
<li></li>
</ol>

**Note:**

1. The following script generates ssh keys in this directory (${HOME}/gcp_keys/).
2. Script logs are saved to a log file in /{HOME}/logs folder.
3. The script will create a DNS zone in cloud DNS. However the nameserver needs to be updated in the DNS of the domain provider.
4. Terraform generates the ansile host file under the directory /terraform/inventory .This file is used to run ansible-playbooks


## Deployment Steps

### Creating a OKD Cluster

1. Script variables can be updated in the /script/variables.sh file. The list of variables are define below

```
SSH_PATH=${HOME}/gcp_keys/id_rsa
SSH_PUB_FILE=${SSH_PATH}.pub
SSH_USER=avinashjha
REGION=europe-west1
ZONE=europe-west1-b
LOG_PATH="${HOME}/logs-$(date +'%Y-%m-%d-%H:%M:%S')"
LOG_FILE="${LOG_PATH}/okd-logs-$(date +'%Y-%m-%d-%H:%M:%S')"
```

2. Terraform variables can be updated in the terraform.tfvars file

--TODO-- 
Realising we can set all the varibles for terraform from the bash script itself. 
--TODO-- 

### Deleting a OKD Cluster

1. Delete the underlying terrafom resources 

```
terraform destroy  -var="gce_ssh_pub_key_file=${HOME}/gcp_keys/id_rsa.pub"
```
