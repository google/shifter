<<<<<<< HEAD
=======
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

>>>>>>> 1da53024b4fc82b48b5a883eeddbe62abf296ed0
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
