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

resource "google_dns_managed_zone" "okd-gcp-zone" {
  project     = module.project.project_id
  name        = "okd-gcp-iac"
  dns_name    = var.dns_master_subdomain
  description = "okd-gcp-iac-zone"
}

resource "google_dns_record_set" "okd-gcp-rs" {
  project = module.project.project_id
  name    = google_dns_managed_zone.okd-gcp-zone.dns_name
  type    = "NS"
  ttl     = 60

  managed_zone = google_dns_managed_zone.okd-gcp-zone.name

  rrdatas = google_dns_managed_zone.okd-gcp-zone.name_servers
}

resource "google_dns_record_set" "os-master-dns-rsa" {
  project = module.project.project_id
  name    = "console.${google_dns_managed_zone.okd-gcp-zone.dns_name}"
  type    = "A"
  ttl     = 300

  managed_zone = google_dns_managed_zone.okd-gcp-zone.name

  rrdatas = [google_compute_address.os-master-addr.address]
}
