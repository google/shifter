
## Introduction

Following would allow us create a OKD Cluster in the GCP Project

## Requirements
<ol>
<li>You have owner permission and you have permission to create service account keys for your GCP project.</li>
<li>OC cli, gcloud cli, terraform is installed.</li>
<li>scripts were tested in a linux machine.</li>
<li></li>
</ol>

**Note:**

1. The following script creates a service account key in order to create a OKD Cluster. After successfull creation of the cluster you can delete/disable the key and create/enable the key again when used again later. Location for the service account key would be `01-projectsetup/sa-keys`.
2. terraform state files and other okd cluster logs are currently stored in the local filesystem, however these can be configured(via providers.tf or other mechanism) to be stored in the GCS bucket or any other persistent file system.

## Folder Structure

1. This repository help you create a openshift 4.10 and openshift 4.9 cluster on GCP. This can further be extended to support other Openshift version(4.x onwards). To start with the binaries for okd 4.10 and okd 4.9 are already kept at `01-projectsetup/okd-installer`. This can be replaced with other newer patched version(if required).
2. `01-projectsetup` contains the scripts to setup pre-requisite for the desired GCP project like modifying org policies, create a service account and creating a public hub zone.
3. `02-appdeployment` contains the manifests for the `bank of anthos` application which will be deployed to the okd cluster.

## Deployment Steps

### Creating a OKD Cluster

1. update the variables in `install.sh` under the MANDATORY_VARIABLES section. Following is a snippet of these variables along with the sample values that you should replace it with.
```
PROJECT_ID=""                    #e.g. : "pm-okd-11"
CLUSTER_NAME=""                  #e.g. : "okd-41"
OKD_VERSION=""                   #e.g. : "4.10"
BILLING_ACCOUNT_ID=""            #e.g. : "0090FE-ED3D81-AF8E3B"
PARENT=""                        #e.g. : "organizations/384628256961"
DOMAIN=""                        #e.g. : "pm-gcp.com."
SSH_KEY_PATH=""                  #e.g. : usr/local/google/home/parasmamgain/.ssh/id_ed25519.pub
# More details on redhat pull secret can be found here https://console.redhat.com/openshift/install/pull-secret
REDHAT_PULL_SECRET='{"auths":{"fake":{"auth":"aWQ6cGFzcwo="}}}'
PROJECT_CREATE="false"           #make this as true if you want to create a new project under the PARENT
```

2. The install.sh would allow you to create multiple clusters in the same project or multiple cluster in different projects.
3. Whenever a public zone is created, we must ensure that the `registrar setup` for this public zone has been performed correctly. This will be one time effort if we plan to use singe project for multiple okd cluster. https://cloud.google.com/dns/docs/update-name-servers .

### Deleting a OKD Cluster

1. update the variables in `destroy.sh` under the MANDATORY_VARIABLES section. Make sure that these variables resembles to the cluster that you have permissions for and you want to delete.Following is a snippet of these variables along with the sample values that you should replace it with.
```
PROJECT_ID=""     # e.g. "pm-okd-11"
CLUSTER_NAME=""   # e.g."okd-41"
OKD_VERSION=""    # e.g."4.10"
```

2. The `destroy.sh` would allow you to delete cluster.
3. The script uses the same service account key which was created by the `install.sh` script, incase the `service-account-key.json` was deleted manually previously then make sure that a valid key exists in the same folder before we can delete the cluster.
Following is what a key creation command may look like:
```
   gcloud iam service-accounts keys create ${SA_JSON_FILENAME} --iam-account=okd-sa@${PROJECT_ID}.iam.gserviceaccount.com

   #gcloud iam service-accounts keys create service-account-key.json --iam-account=okd-sa@pm-okd-11.iam.gserviceaccount.com

   mkdir ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/

   mv ${CWD_PATH}/${SA_JSON_FILENAME} ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/

```

## Output

After the successfull execution of the `install.sh`, the cluster details like endpoint, username and password will be shared via console logs or you can the `okd-cluster/4.x/install-config/<PROJECT_ID>/<CLUSTER_NAME>` for the logs and credentials including the KUBECONFIG for the okd cluster.
