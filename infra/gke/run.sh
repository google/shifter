#! /bin/bash

######################## Optional Vairables(modification not required) ############################
CWD_PATH="$(pwd)"
VALUES_FILE_NAME="development.values"
VALUES_FILE_PATH="$CWD_PATH/$VALUES_FILE_NAME"
#####################################################################
function usage()
{   
    echo "\n"
    echo "The script is used to perform multiple operations on cluster that includes setup and destroy with the help of values file."
    echo "It takes option arguments and process them in given order! "
    echo ""
    echo "USAGE: "
    echo "./run.sh"
    echo "\t-h --help"
    echo "\t--values=development.values (default)"
    echo "\t--setup"
    echo "\t--destroy"
    echo ""
    echo "Example:"
    echo "sh run.sh --values test.values --setup --destroy"
}

function loadvars()
{
    echo "Loading values file: ${VALUES_FILE_PATH}"
    set -o allexport
    source  "$VALUES_FILE_PATH"
    set +o allexport
}

function destroy()
{
   echo "############################################################"
   echo "Initiating the destroy process...." 
   echo "############################################################"

   terraform -chdir=terraform destroy -var-file "$VALUES_FILE_PATH" -auto-approve
}

function setup()
{   
   echo "############################################################"
   echo "Initiating the setup process...." 
   echo "############################################################"

   # Creates pre-reqs for the cluster
   terraform -chdir=terraform init
   terraform -chdir=terraform plan -var-file "$VALUES_FILE_PATH"
   terraform -chdir=terraform apply -var-file "$VALUES_FILE_PATH" -auto-approve

   echo "############################################################"
   echo "Waiting for  60 seconds for resources to be ready..."
   echo "############################################################"
   sleep 60

   ## export kubeconfig on local
   gcloud container clusters get-credentials ${gke_cluster_name} --zone ${gke_location}
}



## we want at least one parameter 
if [ $# -eq 0 ]; then
	usage >&2
	exit 1
fi

## handle shell options here
while [ "$1" != "" ]; do
    PARAM=`echo $1 | awk -F= '{print $1}'`
    VALUE=`echo $1 | awk -F= '{print $2}'`
    case $PARAM in
        -h | --help)
            usage
            exit
            ;;
        --values)
            VALUES_FILE_NAME=$VALUE
            VALUES_FILE_PATH="$CWD_PATH/$VALUES_FILE_NAME"
            loadvars
            ;;    
        --setup)
            setup 
            ;;
        --destroy)
            destroy
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            usage
            exit 1
            ;;
    esac
    shift
done
