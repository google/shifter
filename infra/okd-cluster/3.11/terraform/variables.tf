/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

variable "prefix" {
  type        = string
  description = "Prefix for all resources"
}

variable "master_count" {
  type        = number
  description = "Number of control plane nodes"
  default     = 1
}

variable "node_count" {
  type        = number
  description = "Number of control plane nodes"
  default     = 1
}

variable "project_id" {
  type        = string
  description = "The project id to deploy to"
}

variable "region" {
  type        = string
  description = "The region to deploy the resources into"
}

// project variable

variable "org_id" {
  description = "The organization id for the associated services"
  type        = string
}

variable "folder_id" {
  description = "The folder id where project will be created"
  type        = string
}

variable "billing_account" {
  description = "The ID of the billing account to associated this project with"
  type        = string
}

variable "project_name" {
  description = "The name of the GCP project. Max 16 characters with 3 character business unit code."
  type        = string
}

variable "application_name" {
  description = "The name of application where GCP resources relate"
  type        = string
}

variable "billing_code" {
  description = "The code that's used to provide chargeback information"
  type        = string
}

variable "primary_contact" {
  description = "The primary email contact for the project"
  type        = string
}

variable "secondary_contact" {
  description = "The secondary email contact for the project"
  type        = string
  default     = ""
}

variable "activate_apis" {
  description = "The api to activate for the GCP project"
  type        = list(string)
  default     = []
}

variable "environment" {
  description = "The environment the single project belongs to"
  type        = string
  default     =  ""
}

variable "business_code" {
  description = "Business Code"
  type        = string
  default     =  ""
}

variable "vpc_type" {
  description = "The type of VPC to attach the project to. Possible options are base or restricted."
  type        = string
  default     = ""
}

variable "gce_ssh_user" {
  description = "The ssh user to ssh into the vms"
  type = string
  default = ""
}

variable "gce_ssh_pub_key_file" {
  description = "The path of the ssh pub file attached to the vms for passwordless access"
  type = string
}

variable "master_subdomain" {
  description = "Master domain path"
  type = string

}

variable "public_subdomain" {
  description = "public domain path"
  type = string
}

variable "dns_master_subdomain" {
  description = "Master subdomain entry"
  type = string
}

variable "ssh_user"{
  description = "ssh user"
  type= string
}