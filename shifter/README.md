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

Shifter has extensible methods for inputs and generators.

---

## Processor

Processors are the converts from openshift to kubernetes.

---

## Inputs

Inputs are readers for your existing Openshift application deployment methods

Currently supported inputs:

- **Yaml**

  Yaml input takes a standard OpenShift yaml file and changes certain api calls from OpenShfit specific to standard Kubernetes example: DeploymentConfig to Deployment

- **Templates**

  Template converter takes a Openshift template, converts it into kubernetes compatible resources and outputs given the format required.

- **Cluster**

  Cluster converter takes the resources deployed to a Openshift Namespace, converts those resources into kubernetes compatible resources and outputs given the format required.

---

## Generators

Generators create new code based on your input to be used by standard Kubernetes distributions.

Currently supported generators:

- **Helm**

  Helm charts support the ability to create reusable charts that take input, this is a good fit from moving from Openshift Templates.

- **Yaml**

  Create a standard yaml file for deployment, good for one off deployments such as inputting from yaml.

If you are interested in contributing, see [DEVELOPMENT.md](./DEVELOPMENT.md)

## Converter Usage

### Flags

```
shifter convert
    -f --source-path Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) for Source Files to be Written
    -i --input-format Input format. One of yaml|template (Default: yaml)
    -o --output-path Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written
    -t --output-format Output format (generator to use) One of: yaml|helm
```

#### Processor flags

Processor flags allow you to make changes to the way the processor handles certain objects.

This is achieved using key value pairs passed into the `--pflags` flag.

`--pflags ingress-facing=internal` causes the processor to add the annotation to each ingress object to use a internal load balancer
`--pflags --image-repo=my://registry/address` allows you to change the image registry prefix for your source images

You can chain multiple flags together example:
``go run . convert -t helm -i template -f ./_test/os-nginx-template.yaml -o ./out/helm --pflags ingress-facing=internal --pflags image-repo=gcs://shifter-lz-002``

### Converter Examples:

### Yaml converter

- Running with Local Storage

  `./shifter convert --input-format yaml --source-path ./input.yaml --output-path ./output --output-format yaml `

- Running with GCP Bucket

  `./shifter convert --input-format yaml --source-path gs://bucket/path --output-path gs://bucket/path --output-format yaml `

### Template converter

- Running with Local Storage

  `./shifter convert --input-format template --source-path ./myapp/template.yaml --output-path ./output --output-format helm `

- Running with GCP Bucket

  `./shifter convert --input-format template --source-path gs://bucket/path --output-path gs://bucket/path --output-format helm `

---

## Shifter Server

Shifter also contains a under development Rest API Sever.

## Server Usage

### Flags

```
shifter server
    -p --port Server Port. Default: 8080
    -a --host-address Server Address. Default: 0.0.0.0

    -f --source-path Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) for Source Files to be Written
    -o --output-path Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written
```

### Server Examples:

- Running with Local Storage

  `./shifter server --port 8080 --source-path ./data/source --output-path ./data/output `

- Running with GCP Bucket

  `./shifter server --port 8080 --source-path gs://bucket/source --output-path gs://bucket/output `
