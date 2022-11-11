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

data "template_file" "inventory" {
  template = "${file("${path.cwd}/okd-template/ansible-hosts.template.txt")}"
  vars = {
    master_subdomain = "${var.master_subdomain}"
    public_hostname = "${var.public_subdomain}"
    # master1_hostname = "${google_compute_instance.master1.network_interface.0.network_ip}"
    # infra1_hostname = "${google_compute_instance.infra1.network_interface.0.network_ip}"
    # worker1_hostname = "${google_compute_instance.worker1.network_interface.0.network_ip}"
    # demo_htpasswd = "${var.htpasswd}"
    gcp_project = "${module.project.project_id}"
    ssh_user = "${var.ssh_user}"
  }
}
resource "local_file" "inventory" {
  content     = "${data.template_file.inventory.rendered}"
  filename = "${path.cwd}/inventory/ansible-hosts"
}