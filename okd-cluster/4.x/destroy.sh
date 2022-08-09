#! /bin/bash

###################### Mandatory Vairables ##########################
PROJECT_ID=""     # e.g. "pm-okd-11"
CLUSTER_NAME=""   # e.g."okd-41"
OKD_VERSION=""    # e.g."4.10"

######################## Other Vairables ############################
CWD_PATH="$(pwd)"
echo $CWD_PATH
SA_JSON_FILENAME="service-account-key.json"
PROJECTID_LIST='["'${PROJECT_ID}'"]'


#gcloud iam service-accounts enable okd-sa@${PROJECT_ID}.iam.gserviceaccount.com
gcloud auth activate-service-account okd-sa@${PROJECT_ID}.iam.gserviceaccount.com --key-file=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}
export GOOGLE_APPLICATION_CREDENTIALS=${CWD_PATH}/01-projectsetup/sa-keys/${PROJECT_ID}/${SA_JSON_FILENAME}

echo "############################################################"
echo "Deleting the okd cluster:${CLUSTER_NAME} in project ${PROJECT_ID} ..."
echo "############################################################"
${CWD_PATH}/01-projectsetup/okd-installer/${OKD_VERSION}/openshift-install destroy cluster --log-level=info --dir=${CWD_PATH}/install-config/$PROJECT_ID/$CLUSTER_NAME
