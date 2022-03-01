```
   _____ __    _ ______           
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /    
/____/_/ /_/_/_/  \__/\___/_/     
                                 
----------------------------------------

Openshift to Kubernetes converter
```

# Shifter - Openshift to Kubernetes converter

Easily and quickly convert your RedHat OpenShift workloads to standard kubernetes for Anthos/GKE 

Shifter has extensible methods for inputs and generators.

-----------------
## Convert

### Processor
Processors contains the logic to convert the input object (e.g. Openshift DeploymentConfig) and returns a kubernetes spec for the equivalent standard kubernetes resource (in this case a Deployment). 

### Inputs
Inputs handle the reading in from a filesystem or cluster the files or specs from your current source system

Currently supported inputs:

* **Yaml**

  Yaml input takes a standard OpenShift yaml file(s) which can be one large file or multiple yaml files.
  Yaml files directly describe the resources to provision and don't often offer any templating or parameterization.

* **Templates**

  Templates are a set of parameterized objects which can contain any resource such as Deployments or Services. 
  Parameters are usually provided at the tail of the template file called Parameters and at apply time replace the values in the main body with the values provided.

  Helm charts or Kustomize scripts are a good replacement for templates as both of these provide templating mechanisms.

  Templates are provided in eiter yaml, json or directly deployed to the cluster via the web interface. 


* **Cluster**

  Cluster converter takes the resources deployed to a Openshift Namespace, converts those resources into kubernetes compatible resources and outputs given the format required. 

### Generators

Generators take the converted resources from the processor and creates deployment manifests in the desired format.

Generators read in standard kubernetes objects as a slice array and it's the job of the generator to iterate over and write them out to the desired format.

Currently supported generators:

* **Helm**

  Helm charts support the ability to create reusable charts that take input, this is a good fit from moving from Openshift Templates.

* **Yaml** 

  Create a standard yaml file for deployment, good for one off deployments such as inputting from yaml.

If you are interested in contributing, see [DEVELOPMENT.md](./DEVELOPMENT.md)


-----------------

## Usage

### Flags
```
-f --filename path to either a input file or directory (if reading multiple files)
-i --input-format Input format. One of yaml|template (Default: yaml)
-o --output-path Relative or full path to save the output, if you specify a .yaml or .yml file it will create a multi-document file with all resources, if you specify a directory it will create multiple files per resource type
-t --output-format Output format (generator to use) One of: yaml|helm
```

### Examples:

#### Yaml converter
```./shifter convert --input-format yaml --filename ./input.yaml --output-path ./output --output-format yaml```

### Template converter
```./shifter convert --intput-format template --filename ./myapp/template.yaml --output-path ./output --output-format helm```
