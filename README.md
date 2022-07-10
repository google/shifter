```
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

```

# Openshift to Kubernetes converter

Shifter is a tool which accelerates the migration from OpenShift 3.x / 4.x by translating the applications for Kubernetes, GKE & Anthos and supports migrating to Service Mesh with ASM + Istio Support.
             
## Get Shifting

### Capabilities

1.  Convert existing manifest files from OpenShift to Kubernetes.
2.  Convert or extract manifest files from a running OpenShift cluster.
3.  Convert OpenShift routes/networking to Google ILB/ELB or Istio/ASM virtual services + gateway creation.
4.  Convert OpenShift templates to helm charts.
5.  Convert ImageStreams to Images + Modify on the fly the Container Registry source.
6.  Run locally via a CLI tool or deploy a web-based user interfact.
7.  Use GCS Buckets as the source/destination.

### Components

Shifter has two main components:

##### shifter-ui


##### shifter 


### Run The Latest Deployment Locally 

```

git clone https://github.com/google/shifter 
cd shifter
docker-compose -f docker-compose.yml up

```

### Run Your Development Version Locally 

1) Get the Source

```
git clone https://github.com/google/shifter 
cd shifter
```

2) Then  make your code changes, modifications and add value.

3) Then Build and Run 

```
sh build-deploy-local.sh
```
