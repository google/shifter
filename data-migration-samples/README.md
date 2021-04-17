# Data Migration Samples
There are two shell scripts to demonstrate PV migration flow. 

## Pre-requisite
Please prepare two GKE clusters using [this guide](https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/gce-pd-csi-driver). Be sure to have this option, --addons=GcePersistentDiskCsiDriver, enabled. Once the clusters are created, using the following command to check if there is a standard-rwo storageclass with pd.csi.storage.gke.io as its PROVISIONER. 
<pre>
kubectl get sc
NAME                 PROVISIONER             RECLAIMPOLICY   VOLUMEBINDINGMODE      ALLOWVOLUMEEXPANSION   AGE
premium-rwo          pd.csi.storage.gke.io   Delete          WaitForFirstConsumer   true                   92m
standard (default)   kubernetes.io/gce-pd    Delete          Immediate              true                   92m
<b>standard-rwo         pd.csi.storage.gke.io   Delete          WaitForFirstConsumer   true                   92m</b>
</pre>
If this storageclass is available in both GKE-cluster, we are ready to go.

## s1_remount.sh
This script will step-by-step demonstrate how to detach-and-mount a PV into a new GKE cluster. The detailed steps are explained below:  
### Pre-requisite
1. Copy and paste both clusters context into context1 and context2.
### Execute s1-remount.sh
1. A pod-s with a PVC will be created in the test namespace of context1 cluster. 
2. A file will be copy into the persistent volume for further validation.
3. Patch the persistent volume to have persistentVolumeReclaimPolicy=Retain.
4. Delete pod-s and its corresponding pvc
5. Switch to second cluster and create the same namespace
6. Create a PV to utilize the detached persistent volume
7. Create a new pod-d with pvc, which refers to the PV in step 6. 
8. Validate the file, copied from step-2, is successfully mounted to pod-d. 
9. Clean both environments...

## s2-volumesnapshot.sh
This script demonstrates how to use VolumeSnapshot mechanism to snapshot persistent volume with CSI driver; then detach-and-mount the snapshot onto the target cluster.
### Pre-requisite
1. Copy and paste both clusters context into context1 and context2.
### Execute s2-volumesnapshot.sh
1. Switch to source cluster.
2. Create a pod with persistent disk and then copy an yaml file into the pod for further validation.
3. Specify VolumeSnapshotClass and then perform VolumeSnapshot on the source cluster.
4. Wait the cloned persistent volume ready and then patch it to have persistentVolumeReclaimPolicy="Retain" and then .  
5. Switch to target cluster
6. Create a PV to utilize the detached persistent volume
7. Create a new pod-d with pvc, which refers to the PV in step 7.
8. Validate the file, copied from step 3, is successfully mounted to pod-d.
9. Clean both environments...

## s3-velero.sh
Demonstrate how to use OSS velero to backup/restore persistent volume between two separate clusters.    
### Pre-requisite
1. Copy and paste both clusters context into context1 and context2.
2. Provide desired bucket-name.
3. Provide the file path to store keyfile of the service account. 
### Execute s3-velero.sh
1. Create corresponding GCS (Google Cloud Storage) bucket and build the correponding service account with sufficient privilege. Then, generate the service account key into the filepath, specified by SERVICE_KEY_FILE.
2. Install Velero in both source and target clusters.
---
3. Switch to source cluster.
4. Create pod with Persistent Volume. Then, copy an yaml file into the pod for later checking.
5. Perform velero backup in the source cluster.
6. Switch to the target cluster and perform restore action. 
7. Check if the restored pod contains the file, copied into the pod in step 4.
---
8. Clean all the environments.
