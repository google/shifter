# Shifter Core 
## Openshift to Kubernetes Migration Accelerator

Shifter Core is a tool which accelerates the migration from OpenShift 3.x / 4.x by translating the applications for Kubernetes, GKE & Anthos and supports migrating to Service Mesh with ASM + Istio Support.

<p float="left">
	<img src="../assets/logo.png" alt="shifter logo" />
</p>

Shifter has extensible methods for inputs and generators.

</br>

# Shifter Core CLI

## Shifter
</br>

## Shifter Cluster
Connect to a running OpenShift cluster and interacte with OpenShift objects.
#### Usage 
```
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN <ACTION>
```

#### Flags 
```
-e --cluster-endpoint   OpenShift cluster endpoint                [Required]
-t --token              OpenShift cluster authentication token    [Required]
```

#### Examples 
```
export CLUSTER_ENDPOINT="https://console.okd.<CLUSTER_DOMAIN>:8443"
export BEARER_TOKEN="<BEARER_TOKEN>"

// List All OpenShift Cluster Resources from All OpenShift Namespaces
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN list --all-namespaces

// Export All OpenShift Cluster Resources from All OpenShift Namespaces
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN export --all-namespaces ./output/directory/path
```
</br>

## Shifter Cluster - Convert
Convert takes all the resources from a OpenShift cluster endpoint and converts them to the desired output format on your local disk or GCS bucket
#### Usage 
```
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN convert
```

#### Flags 
```
-e --cluster-endpoint   OpenShift cluster endpoint                [Required]
-t --token              OpenShift cluster authentication token    [Required]
-n --namespace          OpenShift cluster namespace               
-n --all-namespaces     OpenShift all Namespaces or Projects      
-o --output-format      Output format (YAML/HELM)                 [Required]
```
#### Notes

- Supported output-formats include (YAML & HELM)

#### Examples 
```
export CLUSTER_ENDPOINT="https://console.okd.<CLUSTER_DOMAIN>:8443"
export BEARER_TOKEN="<BEARER_TOKEN>"

// Convert All OpenShift Cluster Resources from All OpenShift Namespaces
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN convert --all-namespaces --output-format <FORMAT> ./output/directory/path
```
</br>


## Shifter Cluster - Export
Export takes the resources 'as-is' from a OpenShift cluster endpoint and exports them in yaml format so the manifests can be fed into the shifter conversion process.
#### Usage 
```
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN export
```

#### Flags 
```
-e --cluster-endpoint   OpenShift cluster endpoint                [Required]
-t --token              OpenShift cluster authentication token    [Required]
-n --namespace          OpenShift cluster namespace               
-n --all-namespaces     OpenShift all Namespaces or Projects      
```

#### Examples 
```
export CLUSTER_ENDPOINT="https://console.okd.<CLUSTER_DOMAIN>:8443"
export BEARER_TOKEN="<BEARER_TOKEN>"

// Export All OpenShift Cluster Resources from All OpenShift Namespaces
shifter cluster -e $CLUSTER_ENDPOINT -t $BEARER_TOKEN export --all-namespaces ./output/directory/path
```
</br>

## Shifter Server
Run Shifter Core in Server mode. Allows users to execute API calls to Shifter Core and convert supplied or select objects. In addition to this Shifter Core running in server mode is the backend for Shifter UI.
#### Usage 
```
shifter server 
```

#### Flags 
```
-p --port             Shifter Server Running Port           [Default: 8080   ]
-a --host-address     Shifter Server Running Host Address   [Default: 0.0.0.0]
-f --source-path      Source Relative Local Path               
-o --output-path      Output Relative Local Path      
```

#### Examples 
```
// Run Shifter Core in Server Mode with Defaults
shifter server

// Run Shifter Core on Port 9090
shifter server -p 9090
```
</br>

## Shifter Version
Show the version number of Shifter Core that is currently being executed.
#### Usage 
```
shifter version 
```

#### Flags 
None

#### Examples 
```
// Output Shifter Core's Version
shifter version
```
</br>

## Shifter Help
</br>

</br>

# Shifter Core Components

## Processor

The processor is the component that converts the OpenShift objects to GKE/Anthos compatible objects. 

In OpenShift the following objects are some of the custom resources available to OpenShift and not to other distributions of Kubernetes:

* Projects
* Templates
* DeploymentConfigs
* Routes
* Builds
* ImageStreams

The processor takes the specification of these objects and converts them into the best fit object in GKE/Anthos

* Projects -> Merged with Namesapces
* Templates -> Deployment, Helm Chart
* DeploymentConfigs -> Deployment
* Routes -> Ingress, Internal Load Balancer, External Load Balancer, Istio/ASM VirtualService
* Builds -> CloudBuild manifest
* ImageStream -> Image

The processor is extensible to support further objects, we have a roadmap of items to support.

## Input

Inputs are readers for your existing Openshift application deployment methods,

Currently supported inputs:

- **Yaml**

  Yaml input takes a standard OpenShift yaml manifest file that you have stored on a filesystem or GCS bucket.

- **Templates**

  Template converter takes a Openshift template yaml file, templates can be converted to a templated output format such as Helm Charts.

## Generator

Generators create the resulting output, shifter has been designed so that additional generators or outputters can be created. 

Currently supported generators:

- **Helm**

  Helm charts support the ability to create reusable charts that take input, this is a good fit from moving from Openshift Templates.

- **Yaml**

  Create a standard yaml file for deployment, good for one off deployments such as inputting from yaml.

## Processor flags

Processor flags allow you to make changes to the way the processor handles certain objects.

This is achieved using key value pairs passed into the `--pflags` flag.

`--pflags ingress-facing=internal` causes the processor to add the annotation to each ingress object to use a internal load balancer

`--pflags --image-repo=my://registry/address` allows you to change the image registry prefix for your source images


You can chain multiple flags together example:

``shifter convert -t helm -i template -f ./_test/os-nginx-template.yaml -o ./out/helm --pflags ingress-facing=internal --pflags image-repo=gcs://shifter-lz-002``

## CLI Usage

```
shifter convert
    -f --source-path Relative Local Path (./data/source) or Google Cloud Storage Bucket Path (gs://XXXXXXX/source/) for Source Files to be Written
    -i --input-format Input format. One of yaml|template (Default: yaml)
    -o --output-path Relative Local Path (./data/output) or Google Cloud Storage Bucket Path (gs://XXXXXXX/output/) for Converted Files to be Written
    -t --output-format Output format (generator to use) One of: yaml|helm
```


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