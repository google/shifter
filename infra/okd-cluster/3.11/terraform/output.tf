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