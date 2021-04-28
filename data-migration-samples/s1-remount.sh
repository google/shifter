#!/bin/bash
export namespace="test"
export storageclass="standard-rwo"
export context1="gke_shawn-demo-2021_asia-east1-a_cluster-mlops"
export context2="gke_shawn-demo-2021_asia-east1-a_cluster-security"
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
  storageClassName: ${storageclass}
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
kubectl wait pod/pod-s --for=condition=Ready

echo "copy create-pod-s.yaml into the storage and display it..."
echo "Press any key to continue..."
read k
kubectl cp create-pod-s.yaml pod-s:/usr/share/nginx/html/
kubectl exec -it pod-s -- /bin/sh -c 'ls -l /usr/share/nginx/html'

echo "Prepare detach pvc and pv by removing the pod..."
echo "Press any key to continue..."
read k

pv_name=$(kubectl get pv | grep ${namespace}/pvc-s | awk '{print $1}')
kubectl patch pv ${pv_name} -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'
kubectl delete -f ./create-pod-s.yaml

echo "Switch cluster to ${context2}"
echo "Press any key to continue..."
read k
kubectl config use-context $context2
kubectl create ns ${namespace}
kubectl config set-context --current --namespace ${namespace}
cat << EOF > create-pod-d.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-disk
spec:
  storageClassName: ${storageclass}
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  claimRef:
    namespace: ${namespace}
    name: pvc-d
  gcePersistentDisk:
    pdName: ${pv_name}
    fsType: ext4
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-d
spec:
  storageClassName: ${storageclass}
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-d
  labels:
    app: pod-d
spec:
  volumes:
    - name: task-pv-storage-new
      persistentVolumeClaim:
        claimName: pvc-d
  containers:
    - name: task-pv-container
      image: nginx
      ports:
        - containerPort: 80
          name: "http-server"
      volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: task-pv-storage-new
EOF
kubectl apply -f ./create-pod-d.yaml
while [[ $(kubectl get pods -l app=pod-d -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do 
  echo "waiting for pod-d" && sleep 1; 
done
kubectl exec -it pod-d -- /bin/sh -c 'ls -l /usr/share/nginx/html'

echo "Validating everything is as expected...clean the namespace"
echo "Press any key to continue..."
read k
kubectl delete -f ./create-pod-d.yaml
kubectl delete ns ${namespace}
kubectl config set-context --current --namespace default

echo "Validating everything is as expected...clean the namespace"
echo "Press any key to continue..."
read k
kubectl config use-context ${context1}
kubectl delete ns $namespace
gcloud compute disks delete --quiet ${pv_name}
kubectl config set-context --current --namespace default
