# Openshift to Kubernetes converter

Easily and quickly convert your RedHat Openshift workloads to standard kubernetes for clusters such GKE

Shifter supports two main conversion:

* Yaml converter
  Yaml converter takes a standard OpenShift yaml file and changes certain api calls from OpenShfit specific to standard Kubernetes example: DeploymentConfig to Deployment

* Template converter
  Template converter takes a Openshift template and generates a helm chart that can be deployed against standard kubernetes clusters such as GKE.  This converts certain template types such as DeploymentConfig to Deployment

If you are interested in contributing, see [DEVELOPMENT.md](./DEVELOPMENT.md)

## Usage

### Yaml converter
```./shifter yaml -i input.yaml -o output.yaml```

### Template converter
```./shifter template -i input.yaml -o <directory>```
