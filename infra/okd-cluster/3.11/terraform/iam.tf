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

resource "google_service_account" "os-service" {
  project      = module.project.project_id
  account_id   = "${var.prefix}-os-svc"
  display_name = "Openshift Cluster"
}

resource "google_project_iam_member" "service_account_log_writer" {
  project = module.project.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.os-service.email}"
}

resource "google_project_iam_member" "service_account_metric_writer" {
  project = module.project.project_id
  role    = "roles/monitoring.metricWriter"
  member  = "serviceAccount:${google_service_account.os-service.email}"
}

resource "google_project_iam_member" "service_account_monitoring_viewer" {
  project = module.project.project_id
  role    = "roles/monitoring.viewer"
  member  = "serviceAccount:${google_service_account.os-service.email}"
}

resource "google_project_iam_member" "service_account_owner" {
  project = module.project.project_id
  role    = "roles/owner"
  member  = "serviceAccount:${google_service_account.os-service.email}"
}
