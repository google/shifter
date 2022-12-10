variable "project_id" {
  description = "The GCP project you want to enable APIs and create your project."
  type        = string
}

variable "project_create" {
  description = "This flag will help to either create a new project or use a project that already exists."
  type        = bool
  default     = false
}

variable "vpc_create" {
  description = "This flag will help to either create a new vpc or use a vpc that already exists."
  type        = bool
  default     = true
}

variable "region" {
  type        = string
  default     = "us-central1"
  description = "Region where the cluster and its resources will be created."
}

variable "enable_apis" {
  description = "Whether to actually enable the APIs. If false, this module is a no-op."
  default     = true
  type        = bool
}

variable "disable_services_on_destroy" {
  description = "Whether project services will be disabled when the resources are destroyed. https://www.terraform.io/docs/providers/google/r/google_project_service.html#disable_on_destroy"
  default     = false
  type        = bool
}

variable "disable_dependent_services" {
  description = "Whether services that are enabled and which depend on this service should also be disabled when this service is destroyed. https://www.terraform.io/docs/providers/google/r/google_project_service.html#disable_dependent_services"
  default     = false
  type        = bool
}

variable "vpc_network_name" {
  description = "The vpc network name that needs to be used."
  type        = string
  default     = "gke-vpc"
}

variable "vpc_subnetwork_name" {
  description = "The vpc subnetwork name that needs to be used."
  type        = string
  default     = "gke-development"
}

variable "vpc_subnetwork_cidr" {
  description = "The vpc subnetwork cidr that needs to be used."
  type        = string
  default     = "10.0.0.0/24"
}

variable "gke_pods_secondary_cidr" {
  description = "The vpc subnetwork secondary range cidr that needs to be used for pods in gke."
  type        = string
  default     = "10.57.0.0/17"
}

variable "gke_services_secondary_cidr" {
  description = "The vpc subnetwork secondary range cidr that needs to be used for services in gke."
  type        = string
  default     = "10.57.128.0/22"
}

variable "gke_cluster_name" {
  description = "The gke cluster name that needs to be used."
  type        = string
  default     = "gke-for-okd-workloads"
}

variable "gke_nodepool_name" {
  description = "The gke nodepool name that will be used."
  type        = string
  default     = "default-nodepool"
}

variable "gke_nodepool_count" {
  description = "The gke nodepool node count."
  type        = number
  default     = 1
}
