#!/bin/zsh
export namespace="test"
export namespace2="test"
export context1="gke_shawn-demo-2021_asia-east1-a_cluster-mlops"
export context2="gke_shawn-demo-2021_asia-east1-a_cluster-security"
export BUCKET="velero-backup-demo"
export SERVICE_ACCOUNT_EMAIL="velero@shawn-demo-2021.iam.gserviceaccount.com"
export SERVICE_KEY_FILE="${HOME}/workspace/gcp-keys/shawn-demo-2021-velero.key"
export PROJECT_ID="shawn-demo-2021"

function pause_with_msg() {
  msg=$1
  echo "${msg}"
  echo "Press any key to continue..."
  read -k
}
function create_sa() {
  ROLE_PERMISSIONS=(
    compute.disks.get
    compute.disks.create
    compute.disks.createSnapshot
    compute.snapshots.get
    compute.snapshots.create
    compute.snapshots.useReadOnly
    compute.snapshots.delete
    compute.zones.get
  )

  gcloud iam roles create velero.server \
    --project $PROJECT_ID \
    --title "Velero Server" \
    --permissions "$(IFS=","; echo "${ROLE_PERMISSIONS[*]}")"

  gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:$SERVICE_ACCOUNT_EMAIL \
    --role projects/$PROJECT_ID/roles/velero.server

  gsutil iam ch serviceAccount:$SERVICE_ACCOUNT_EMAIL:objectAdmin gs://${BUCKET}
  gcloud iam service-accounts keys create ${SERVICE_KEY_FILE} \
    --iam-account $SERVICE_ACCOUNT_EMAIL
}

function install_velero() {
  pause_with_msg "Prepare velero install on both clusters"
  velero install --provider gcp --plugins velero/velero-plugin-for-gcp:v1.2.0 --bucket $BUCKET --use-restic  --secret-file ${SERVICE_KEY_FILE} --kubecontext ${context1}
  velero install --provider gcp --plugins velero/velero-plugin-for-gcp:v1.2.0 --bucket $BUCKET --use-restic  --secret-file ${SERVICE_KEY_FILE} --kubecontext ${context2}
}
function apply_pod() {
  pause_with_msg "Deploy pod with persistent disk in the first cluster..."
  kubectl config use-context ${context1}
  kubectl create ns $namespace
  kubectl config set-context --current --namespace $namespace
cat << EOF > create-pod-s.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-s
  namespace: ${namespace}
spec:
  storageClassName: standard-rwo
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-s
  namespace: ${namespace}
  labels:
    app: pod-s
  annotations:
    backup.velero.io/backup-volumes: task-pv-storage
spec:
  volumes:
    - name: task-pv-storage
      persistentVolumeClaim:
        claimName: pvc-s
  containers:
    - name: nginx-container
      image: nginx
      ports:
        - containerPort: 80
          name: "http-server"
      volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: task-pv-storage
EOF
  kubectl apply -f ./create-pod-s.yaml

  while [[ $(kubectl get pods -l app=pod-s -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for pod" && sleep 1; done
}

function backup_pod() {
  pause_with_msg "Copy a create-pod-s.yaml file into the pod..."
  kubectl cp create-pod-s.yaml pod-s:/usr/share/nginx/html/
  kubectl exec -it pod-s -- /bin/sh -c 'ls -l /usr/share/nginx/html'

  pause_with_msg "Create backup for namespace $namespace"
  velero backup create ${namespace}-bk --include-namespaces ${namespace} --wait
}

function restore_pod() {
  pause_with_msg "Restore at cluster2..."
  kubectl config use-context ${context2}
  kubectl config set-context --current --namespace $namespace
  while true; do
    bk=$(kubectl -n velero get backup | grep ${namespace}-bk)
    if [[ ! -z $bk ]]; then
       break;
    fi
    echo "waiting backup to be ready"
    sleep 5;
  done
  velero restore create $restorename --from-backup ${namespace}-bk --include-namespaces ${namespace} --namespace-mappings ${namespace}:${namespace2}

  pause_with_msg "Check the result..."
  kubectl exec -it pod-s -n ${namespace2} -- /bin/sh -c 'ls -l /usr/share/nginx/html'
}

function delete_pod() {
  pause_with_msg "Delete all pods and velero installation"
  velero backup delete ${namespace}-bk --kubecontext $context1
  kubectl delete ns $namespace --context ${context1}
  kubectl delete ns $namespace2 --context ${context2}
}

function delete_velero() {
  for c in ${context1} ${context2}; do
    kubectl delete namespace/velero clusterrolebinding/velero --context $c
    kubectl delete crds -l component=velero --context $c
  done
}

#create_sa
install_velero
apply_pod
backup_pod
restore_pod
delete_pod
#delete_velero
