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

data "google_compute_zones" "zones" {
  region  = var.region
  project = module.project.project_id
}

data "google_compute_image" "centos7" {
  family  = "centos-7"
  project = "centos-cloud"
}

resource "google_compute_instance_template" "okd-tpl" {
  project     = module.project.project_id
  description = "Template for OKD compute instances"

  tags = ["okd", "openshift"]

  name_prefix  = "${var.prefix}-cluster-"
  machine_type = "e2-standard-4"

  disk {
    boot         = true
    auto_delete  = true
    source_image = data.google_compute_image.centos7.self_link
    disk_type    = "pd-ssd"
    disk_size_gb = "100"
  }

  network_interface {
    subnetwork = module.network.subnetwork.0.self_link
  }

  service_account {
    email  = google_service_account.os-service.email
    scopes = ["cloud-platform"]
  }
  metadata = {
    ssh-keys = "${var.gce_ssh_user}:${file(var.gce_ssh_pub_key_file)}"
  }
  
  metadata_startup_script = <<EOF
yum -y install wget git net-tools bind-utils yum-utils iptables-services bridge-utils bash-completion kexec-tools sos psacct httpd-tools java-1.8.0-openjdk-devel.x86_64
yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
yum -y --enablerepo=epel install ansible pyOpenSSL python-pip python-passlib
yum -y install docker-1.13.1
pip install --upgrade passlib
curl -sSO https://dl.google.com/cloudagents/add-google-cloud-ops-agent-repo.sh
bash add-google-cloud-ops-agent-repo.sh --also-install
EOF
}


resource "google_compute_instance_from_template" "os-control-plane" {
  count   = var.master_count
  project = module.project.project_id
  name    = "${var.prefix}-cluster-cp-${count.index}"
  zone    = data.google_compute_zones.zones.names[0]

  source_instance_template = google_compute_instance_template.okd-tpl.id

  tags = ["okd", "os-cp", "openshift"]
}

resource "google_compute_instance_from_template" "os-data-plane" {
  count   = var.node_count
  project = module.project.project_id
  name    = "${var.prefix}-cluster-dp-${count.index}"
  zone    = data.google_compute_zones.zones.names[0]

  source_instance_template = google_compute_instance_template.okd-tpl.id

  tags = ["okd", "os-dp", "openshift"]
}

resource "google_compute_instance_from_template" "os-infra-plane" {
  count   = 1
  project = module.project.project_id
  name    = "${var.prefix}-cluster-in-${count.index}"
  zone    = data.google_compute_zones.zones.names[0]

  source_instance_template = google_compute_instance_template.okd-tpl.id

  tags = ["okd", "os-cp", "openshift"]
}

resource "google_compute_instance_from_template" "os-deployer" {
  count   = 1
  project = module.project.project_id
  name    = "${var.prefix}-cluster-deployer-${count.index}"
  zone    = data.google_compute_zones.zones.names[0]

  machine_type = "e2-medium"

  source_instance_template = google_compute_instance_template.okd-tpl.id

  tags = ["okd", "deployer", "openshift"]
}
