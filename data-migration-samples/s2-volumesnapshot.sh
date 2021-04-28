#!/bin/bash
export namespace="test"
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

kubectl cp create-pod-s.yaml pod-s:/usr/share/nginx/html/
kubectl exec -it pod-s -- /bin/sh -c 'ls -l /usr/share/nginx/html'
echo "Prepare snapshot..."
echo "Press any key to continue..."
read k

cat << EOF > snapshot-s.yaml
# snapshot-s.yaml
apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshotClass
metadata:
  name: gke-ssc
driver: pd.csi.storage.gke.io
deletionPolicy: Delete
---
#snapshot-example.yaml
apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshot
metadata:
  name: snapshot-s
spec:
  volumeSnapshotClassName: gke-ssc
  source:
    persistentVolumeClaimName: pvc-s
EOF
kubectl apply -f ./snapshot-s.yaml
while [[ $(kubectl get volumesnapshot -o jsonpath='{.items[0].status.readyToUse}') != "true" ]];
  do echo "wait for volumesnapshot" && sleep 1; 
done

kubectl get volumesnapshot \
  -o custom-columns='NAME:.metadata.name,READY:.status.readyToUse'

echo "Clone persistent-volume d"
echo "Press any key to continue..."
read k
cat << EOF > clone-d.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-d
spec:
  dataSource:
    name: snapshot-s
    kind: VolumeSnapshot
    apiGroup: snapshot.storage.k8s.io
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
  name: pod-inter
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
kubectl apply -f ./clone-d.yaml
while [[ ! $(kubectl get pv | grep ${namespace}/pvc-d) ]];
  do echo "wait for cloned pv" && sleep 1; 
done
pv_name=$(kubectl get pv | grep ${namespace}/pvc-d | awk '{print $1}')
kubectl patch pv ${pv_name} -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'

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
  storageClassName: standard-rwo
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

echo "Validating everything is as expected...clean the namespace"
echo "Press any key to continue..."
read k
kubectl config use-context ${context1}
kubectl delete -f ./clone-d.yaml
kubectl delete volumesnapshot snapshot-s
kubectl delete -f ./create-pod-s.yaml
kubectl delete ns $namespace
gcloud compute disks delete --quiet ${pv_name}
