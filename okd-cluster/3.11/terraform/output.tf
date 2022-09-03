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

 output "project_id" {
  description = "Project ID"
  value       = module.project.project_id
}


output "bastion" {
  description = "The hostname of the bastion host"
  value = google_compute_instance_from_template.os-deployer[0].name
}

output "master" {
  description = "The hostname of the master host"
  value = google_compute_instance_from_template.os-control-plane[0].name
}

output "infra" {
  description = "The hostname of the infra host"
  value = google_compute_instance_from_template.os-infra-plane[0].name
}

output "dns-name-server" {
    description  = "The nameserver created "
    value = google_dns_managed_zone.okd-gcp-zone.name_servers
}

output "google_compute_address" {
  description = "Address associated with LB"
  value = google_compute_address.os-master-addr.address
}