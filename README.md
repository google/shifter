```
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
```

# Openshift to Kubernetes converter

Easily and quickly convert your RedHat OpenShift workloads to standard kubernetes for Anthos/GKE



## Get Shifting

### Run The Latest Deployment Locally 

1) Get the Source & Run Latest the latest Docker Release
```

git clone https://github.com/google/shifter 
cd shifter
docker-compose -f docker-compose.yml up

```

2) Open your browser to [http://localhost:9090](http://localhost:9090)
   
   

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

4) Open your browser to [http://localhost:9090](http://localhost:9090)