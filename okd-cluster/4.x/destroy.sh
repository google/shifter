#! /bin/bash

###################### Mandatory Vairables ##########################
PROJECT_ID="pm-singleproject-20"     # e.g. "pm-okd-11"
CLUSTER_NAME="okd41"   # e.g."okd-41"
OKD_VERSION="4.10"    # e.g."4.10"

######################## Other Vairables ############################
CWD_PATH="$(pwd)"
echo $CWD_PATH
SA_JSON_FILENAME="service-account-key.json"
PROJECTID_LIST='["'${PROJECT_ID}'"]'

echo "Getting Secret"

mkdir -p ${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/
gcloud secrets versions access 1 --secret="okd-service-account" --out-file=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}

#gcloud iam service-accounts enable okd-sa@${PROJECT_ID}.iam.gserviceaccount.com
gcloud auth activate-service-account okd-sa@${PROJECT_ID}.iam.gserviceaccount.com --key-file=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}
export GOOGLE_APPLICATION_CREDENTIALS=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}

echo "############################################################"
echo "Cloning artifacts from GCS bucket"
echo "############################################################"

gcloud storage cp gs://shifter-tfstate/builds/plan-file/v0.3.1/* ${CWD_PATH}/install-config/pm-singleproject-20/okd41


echo "############################################################"
echo "Deleting the okd cluster:${CLUSTER_NAME} in project ${PROJECT_ID} ..."
echo "############################################################"
${CWD_PATH}/01-projectsetup/okd-installer/${OKD_VERSION}/openshift-install destroy cluster --log-level=info --dir=${CWD_PATH}/install-config/$PROJECT_ID/$CLUSTER_NAME
