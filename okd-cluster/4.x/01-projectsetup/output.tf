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

