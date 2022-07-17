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
  }
}
resource "local_file" "inventory" {
  content     = "${data.template_file.inventory.rendered}"
  filename = "${path.cwd}/inventory/ansible-hosts"
}