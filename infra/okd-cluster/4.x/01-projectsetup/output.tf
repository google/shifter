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

output "nameservers" {
  description = "NS to be registered with the DNS"
  value       = ({
    for key in var.projectid_list :
      key => [{
        nameserver = module.dns-public-zone[key].name_servers
      }]
    })
}

output "bucket" {
  description = "Bucket name to store state and other artifacts"
  value       = module.gcs-automation
}

output "sa" {
  description = "SA name to use for impersonation and resource creation/deletion/modification."
 value        = module.okd-sa
  sensitive   = true
}

