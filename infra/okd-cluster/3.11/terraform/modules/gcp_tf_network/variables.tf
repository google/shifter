variable "name" {
  type        = string
  description = "The name of the resources to be deployed"
}

variable "prefix" {
  type        = string
  description = "The prefix for all resources deployed in the project"
}

variable "project_id" {
  type        = string
  description = "The project id to deploy the network to"
}

variable "subnets" {
  type        = list(any)
  description = "List of map of subnetworks and the regions to create"
}

variable "default_region" {
  type        = string
  description = "The default region to use"
  default     = "europe-west1"
}
