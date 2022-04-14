```
   _____ __    _ ______
  / ___// /_  (_) __/ /____  _____
  \__ \/ __ \/ / /_/ __/ _ \/ ___/
 ___/ / / / / / __/ /_/  __/ /
/____/_/ /_/_/_/  \__/\___/_/


```

# Openshift to Kubernetes converter

Easily and quickly convert your RedHat OpenShift workloads to standard kubernetes for Anthos/GKE.

Shifter has extensible methods for inputs and generators.



## Processor

Processors are the converts from openshift to kubernetes.



## Supported Inputs

Inputs are readers for your existing Openshift application deployment methods

Currently supported inputs:

- **Yaml**

  Yaml input takes a standard OpenShift yaml file and changes certain api calls from OpenShfit specific to standard Kubernetes example: DeploymentConfig to Deployment

- **Templates**

  Template converter takes a Openshift template, converts it into kubernetes compatible resources and outputs given the format required.

- **Cluster**

  Cluster converter takes the resources deployed to a Openshift Namespace, converts those resources into kubernetes compatible resources and outputs given the format required.

## Generators

Generators create new code based on your input to be used by standard Kubernetes distributions.

Currently supported generators:

- **Helm**

  Helm charts support the ability to create reusable charts that take input, this is a good fit from moving from Openshift Templates.

- **Yaml**

  Create a standard yaml file for deployment, good for one off deployments such as inputting from yaml.



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
Shifter Convert has several configurations options made available via the way of flags. 

| **Flag**       | **Description**        |
| ----------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| -f  <br> --filename | path to either a input file or directory (if reading multiple files).|
| -i  <br> --input-format |Input format. One of yaml|template (Default: yaml).|
| -o <br> --output-path | Relative or full path to save the output, if you specify a .yaml or .yml file it will create a multi-document file with all resources, if you specify a directory it will create multiple files per resource type.|
| -t  <br> --output-format |Output format (generator to use) One of: yaml|helm|




### Example Usage:

| **CLI**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: | YAML Conversion using Go Binary via CLI |
| |`./shifter convert --input-format yaml --filename ./input.yaml --output-path ./output --output-format yaml` |
| Examples #2: | Template Conversion using Go Binary via CLI |
| |`./shifter convert --intput-format template --filename ./myapp/template.yaml --output-path ./output --output-format helm` |


| **Docker Container**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |YAML Conversion in Docker Container with local System Volume Mount |
| |```docker run --rm -it -u $(id -u) -v ${PWD}/[SRC_FILE/DIR_PATH]/:/source images.shifter.cloud/shifter ./shifter convert --input-format yaml --filename "/source/" --output-path "/source/results" --output-format yaml ```|
| Examples #2: |Template Conversion in Docker Container with local System Volume Mount |
| |```docker run --rm -it -u $(id -u) -v ${PWD}/[SRC_FILE/DIR_PATH]/:/source images.shifter.cloud/shifter ./shifter convert --input-format template --filename "/source/" --output-path "/source/results" --output-format helm ``` |

<br><br><br><br>

```
	 _____ __    _ ______            
	/ ___// /_  (_) __/ /____  _____       ___   ____ ____ ______  ___    ___   ____
	\__ \/ __ \/ / /_/ __/ _ \/ ___/      / _ \ / __// __//_  __/ / _ |  / _ \ /  _/
   ___/ / / / / / __/ /_/  __/ /         / , _// _/ _\ \   / /   / __ | / ___/_/ /
  /____/_/ /_/_/_/  \__/\___/_/         /_/|_|/___//___/  /_/   /_/ |_|/_/   /___/ 
                                 
```

## Shifter Server
Shifter Server has several configurations options made available via the way of flags. 

| **Flag**       | |**Description**        |
| ----------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| -p | --port |Server Port: Deafult (8080)|


### Example Usage:

| **CLI**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |Run Shifter Server Binary locally on port "8080" |
| |`./shifter server -p 8080` |


| **Docker Container**       |       |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| Examples #1: |Run Shifter Server as a container on port "8080" |
| |`docker run -it -p 8080:8080 images.shifter.cloud/shifter:latest ./shifter server -p 8080` |



# Contribution

If you are interested in contributing, see [DEVELOPMENT.md](./DEVELOPMENT.md)

---
