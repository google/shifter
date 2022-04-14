
```
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/

----------------------------------------
```

# OpenShift to Kubernetes converter

Easily and quickly convert your RedHat OpenShift workloads to standard Kubernetes for Anthos/GKE 

Shifter has extensible methods for inputs and generators.

---

## Processor

Processors are the converts from OpenShift to Kubernetes.

---

## Supported Inputs

Inputs are readers for your existing Openshift application deployment methods

Currently supported inputs:

- **Yaml**

  Yaml input takes a standard OpenShift yaml file and changes certain api calls from OpenShift specific to standard Kubernetes example: DeploymentConfig to Deployment

- **Templates**

  Template converter takes a OpenShift template, converts it into Kubernetes compatible resources and outputs given the format required.

- **Cluster**

  Cluster converter takes the resources deployed to a OpenShift Namespace, converts those resources into Kubernetes compatible resources and outputs given the format required. 

---

## Generators

Generators create new code based on your input to be used by standard Kubernetes distributions.

Currently supported generators:

- **Helm**

  Helm charts support the ability to create reusable charts that take input, this is a good fit from moving from OpenShift Templates.

- **Yaml**

  Create a standard yaml file for deployment, good for one off deployments such as inputting from yaml.

If you are interested in contributing, see [DEVELOPMENT.md](./DEVELOPMENT.md)

---

<br>
<br>

# Usage

The Shifter CLI can be executed in several modes. Each of which has it's own subroutine.

| Shifter CLI Usage ||
| ----------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------
 |
| **Routine**       | **Description**                                                                                                                                      |
| ./shifter convert | Takes an OpenShift source input format and File/Directory as flags and converts to a specified Output File/Directory in the specified output format. |
| ./shifter server  | Starts Shifter as a HTTP WebServer allowing Shifter functionality to be made available by Rest API Endpoints.                                        |

<br>

## Shifter Convert
---
<br>

| **Flag**       | |**Description**        |
| ----------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| -f | --filename |path to either a input file or directory (if reading multiple files).|
| -i | --input-format |Input format. One of yaml|template (Default: yaml).|
| -o | --output-path |Relative or full path to save the output, if you specify a .yaml or .yml file it will create a multi-document file with all resources, if you specify a directory it will create multiple files per resource type.|
| -t | --output-format |Output format (generator to use) One of: yaml|helm|

<br>
---

### Example Usage:

| **CLI**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: | YAML Conversion |
| |`./shifter convert --input-format yaml --filename ./input.yaml --output-path ./output --output-format yaml` |
| Examples #3: | Template Conversion |
| |`./shifter convert --intput-format template --filename ./myapp/template.yaml --output-path ./output --output-format helm` |


| **Docker Container**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |Run Shifter Server as a container on port "8080" |
| |`docker run -it -p 8080:8080 shifter.cloud/shifter:latest ./shifter server -p 8080` |


<br><br>

## Shifter Server
---
<br>

| **Flag**       | |**Description**        |
| ----------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| -p | --port |Server Port: Deafult (8080)|

<br>
---

### Example Usage:

| **CLI**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |Run Shifter Server Binary locally on port "8080" |
| |`./shifter server -p 8080` |


| **Docker Container**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |Run Shifter Server as a container on port "8080" |
| |`docker run -it -p 8080:8080 shifter.cloud/shifter:latest ./shifter server -p 8080` |

