variable "project_id" {
  description = "The GCP project you want to enable APIs and create your project."
  type        = string
  default     = ""
}

variable "projectid_list" {
  description = "The GCP project you want to enable APIs and create your project."
  type        = list(string)
}

variable "project_create" {
  description = "This flag will help to either create a new project or use a project that already exists."
  type        = bool
  default     = false
}

variable "domain" {
  description = "The domain name owned by the user which will then be used for creating a public facing cluster."
  type = string
}

variable "cluster_name" {
  type = string
  description = "Name of OKD Cluster(max length restricted to 10)"
}

variable "ssh_key_path" {
  type = string
  default = ""
  description = "Path for the ssh key public certificate for the machine which can be used to troubleshoot the clusters."
}

variable "region" {
  type    = string
  default = "us-central1"
  description = "Region where the cluster and its resources will be created."
}

variable "billing_account_id" {
  type = string
  description = "Billing account id to be associated with the project."
}

variable "redhat_pull_secret" {
  type = string
  description = "Redhat pull secret(default it uses a generic pull secret)."
}

variable "parent" {
  type = string
  description = "parent orgId under which the project exists(or will be created)."
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
