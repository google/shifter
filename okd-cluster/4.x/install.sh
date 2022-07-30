#! /bin/bash


###################### Mandatory Vairables ##########################
PROJECT_ID=""                    #e.g. : "pm-okd-11"
CLUSTER_NAME=""                  #e.g. : "okd-41"
OKD_VERSION=""                   #e.g. : "4.10" or "4.9"
BILLING_ACCOUNT_ID=""            #e.g. : "xxxxxx-yyyyyy-zzzzzz"
PARENT=""                        #e.g. : "organizations/111222333444"
DOMAIN=""                        #e.g. : "google.com."
SSH_KEY_PATH=""                  #e.g. : usr/local/google/home/username/.ssh/id_ed25519.pub
# More details on redhat pull secret can be found here https://console.redhat.com/openshift/install/pull-secret
REDHAT_PULL_SECRET='{"auths":{"fake":{"auth":"aWQ6cGFzcwo="}}}'
PROJECT_CREATE="false"           #make this as true if you want to create a new project under the PARENT

######################## Optional Vairables(modification not required) ############################
CWD_PATH="$(pwd)"
echo $CWD_PATH
SA_JSON_FILENAME="service-account-key.json"
PROJECTID_LIST='["'${PROJECT_ID}'"]'
OKD_INSTALLABALE_VERSION=""

###################################################################
echo "############################################################"
echo "Validating Mandatory Variables"
echo "############################################################"
ERROR=""
if [ -z "$PROJECT_ID" ]
then
      ERROR="PROJECT_ID should not be empty"
fi

if [ -z "$CLUSTER_NAME" ]
then
      ERROR="$ERROR""\nCLUSTER_NAME should not be empty"
fi

if [ "$OKD_VERSION" != "4.9" ] && [ "$OKD_VERSION" != "4.10" ]
then
      ERROR="$ERROR""\nOKD_VERSION should be 4.9 or 4.10"
fi

if [ -z "$BILLING_ACCOUNT_ID" ]
then
      ERROR="$ERROR""\nBILLING_ACCOUNT_ID should not be empty"
fi

if [ -z "$PARENT" ]
then
      ERROR="$ERROR""\nPARENT should not be empty"
fi

if [ -z "$DOMAIN" ]
then
      ERROR="$ERROR""\nDOMAIN should not be empty"
fi

if [ -z "$REDHAT_PULL_SECRET" ]
then
      ERROR="$ERROR""\nREDHAT_PULL_SECRET should not be empty"
fi

if [ -z "$PROJECT_CREATE" ]
then
      ERROR="$ERROR""\nPROJECT_CREATE should not be empty"
fi

if [ -z "$ERROR" ]
then
     echo "Basic Validation completed"
else
      echo -e "$ERROR"
      exit
fi

echo "############################################################"
echo "Configuring project and setting up project pre-requisite..."
echo "############################################################"

# Creates pre-reqs for the cluster
terraform -chdir=01-projectsetup init
terraform -chdir=01-projectsetup plan -var "projectid_list=${PROJECTID_LIST}" -var "cluster_name=${CLUSTER_NAME}" -var "billing_account_id=${BILLING_ACCOUNT_ID}" -var "parent=${PARENT}" -var "redhat_pull_secret=${REDHAT_PULL_SECRET}" -var "domain=${DOMAIN}" -var "ssh_key_path=${SSH_KEY_PATH}" -var "project_create=${PROJECT_CREATE}"
terraform -chdir=01-projectsetup apply -var "projectid_list=${PROJECTID_LIST}" -var "cluster_name=${CLUSTER_NAME}" -var "billing_account_id=${BILLING_ACCOUNT_ID}" -var "parent=${PARENT}" -var "redhat_pull_secret=${REDHAT_PULL_SECRET}" -var "domain=${DOMAIN}" -var "ssh_key_path=${SSH_KEY_PATH}" -var "project_create=${PROJECT_CREATE}" --auto-approve

echo "##############################################################################"
echo "Please Ensure if you have updated the Nameservers in your domain name registar"
echo "##############################################################################"
read -p "Press enter to continue"

echo "############################################################"
echo "Waiting for  60 seconds for resources to be ready..."
echo "############################################################"
sleep 60s
# Other versions can be downloaded from https://github.com/openshift/okd/releases/

# Download okd installer and oc cli based on the OKD_VERSION
if [ -f ${CWD_PATH}/01-projectsetup/okd-installer/${OKD_VERSION}/openshift-install ];
then
   echo "############################################################"
   echo "File openshift-install exist."
   echo "############################################################"
else
   if [ ${OKD_VERSION} == "4.10"]
   then
      OKD_INSTALLABALE_VERSION="4.10.0-0.okd-2022-06-10-131327"
   elif [ ${OKD_VERSION} == "4.9" ]
   then
      OKD_INSTALLABALE_VERSION="4.9.0-0.okd-2022-02-12-140851"
   else
      OKD_INSTALLABALE_VERSION="4.10.0-0.okd-2022-06-10-131327"
   fi
   echo "############################################################"
   echo "File openshift-install does not exist. Downloading openshift-install file for OKD+VERSION: ${OKD_INSTALLABALE_VERSION}"
   echo "############################################################"
   wget -O openshift-install-linux.tar.gz https://github.com/openshift/okd/releases/download/${OKD_INSTALLABALE_VERSION}/openshift-install-linux-${OKD_INSTALLABALE_VERSION}.tar.gz
   tar -xvf openshift-install-linux.tar.gz
   chmod +x openshift-install
   mv openshift-install ${CWD_PATH}/01-projectsetup/okd-installer/${OKD_VERSION}/
fi

# Download OC CLI
#wget -O openshift-clientinstall-linux.tar.gz https://github.com/openshift/okd/releases/download/4.10.0-0.okd-2022-06-10-131327/openshift-client-linux-4.10.0-0.okd-2022-06-10-131327.tar.gz
#tar -xvf openshift-clientinstall-linux.tar.gz
#chmod +x oc
#mv oc /usr/bin/local/




#create service account key
if [ -f ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME} ];
then
   echo "############################################################"
   echo "File ${SA_JSON_FILENAME} exist."
   echo "############################################################"
else
   echo "############################################################"
   echo "File ${SA_JSON_FILENAME} does not exist. Creating Service Account key file"
   echo "############################################################"
   gcloud iam service-accounts keys create ${SA_JSON_FILENAME} --iam-account=okd-sa@${PROJECT_ID}.iam.gserviceaccount.com
   #gcloud iam service-accounts keys create test.json --iam-account=okd-sa@pm-okd-11.iam.gserviceaccount.com
   mkdir ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/
   mv ${CWD_PATH}/${SA_JSON_FILENAME} ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/
fi

gcloud config set account okd-sa@${PROJECT_ID}.iam.gserviceaccount.com
#Activates the SA to be used
gcloud auth activate-service-account okd-sa@${PROJECT_ID}.iam.gserviceaccount.com --key-file=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}

#Exports the APPLICATION Credentials to be used bu the openshift installer
export GOOGLE_APPLICATION_CREDENTIALS=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}

echo "#################################################################"
echo "Creating OKD Cluster:${CLUSTER_NAME} in project ${PROJECT_ID} ..."
echo "#################################################################"

#Performs installation of the okd cluster
${CWD_PATH}/01-projectsetup/okd-installer/${OKD_VERSION}/openshift-install create cluster --log-level=info --dir=${CWD_PATH}/install-config/$PROJECT_ID/$CLUSTER_NAME


export USERNAME="kubeadmin"
export PASSWORD=`cat ${CWD_PATH}/install-config/${PROJECT_ID}/${CLUSTER_NAME}/auth/kubeadmin-password`
export KUBECONFIG=${CWD_PATH}/install-config/${PROJECT_ID}/${CLUSTER_NAME}/auth/kubeconfig
#Disable the service account
#gcloud iam service-accounts disable okd-sa@${PROJECT_ID}.iam.gserviceaccount.com

echo "#################################################################"
echo "Deploying Application workload in the cluster:${CLUSTER_NAME} ..."
echo "#################################################################"

#update the providers with appropriate kubeconfig
#sed -e "s|KUBE-CONFIG|$KUBECONFIG|" 02-appdeployment/hello-python/provider.tf.template > 02-appdeployment/hello-python/provider.tf
#  Deploying application in the cluster
## Deploying hello-python flask application in the okd cluster

#terraform -chdir=02-appdeployment/hello-python init
#terraform -chdir=02-appdeployment/hello-python plan
#terraform -chdir=02-appdeployment/hello-python apply --auto-approve

## Deploying bank of anthos modified yaml
# Github URL : https://github.com/GoogleCloudPlatform/bank-of-anthos/blob/main/docs/environments.md#non-gke-kubernetes-clusters

oc apply -f ${CWD_PATH}/02-appdeployment/bank-of-anthos/kubernetes-manifests/jwt/jwt-secret.yaml
oc apply -f ${CWD_PATH}/02-appdeployment/bank-of-anthos/kubernetes-manifests
echo "############################################################"
echo "Waiting for  60 seconds for workloads to be ready..."
echo "############################################################"
sleep 60s
oc get pods

echo "############################################################"
echo "Endpoint"
echo "############################################################"
oc get service frontend | awk '{print $4}'

echo "##################################################################"
echo "Get Token and the Cluster API Endpoint to be used for the shifter"
echo "##################################################################"
oc login --username=$USERNAME --password=$PASSWORD
export TOKEN=$(grep 'token:' $KUBECONFIG | tail -n1); TOKEN=${TOKEN//*token: /};
export CLUSTER_API_ENDPOINT=$(grep 'server:' $KUBECONFIG | tail -n1); CLUSTER_API_ENDPOINT=${CLUSTER_API_ENDPOINT//*server: /};

#echo $TOKEN
#echo $CLUSTER_API_ENDPOINT

echo "##################################################################"
