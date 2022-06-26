output "nameservers" {
  description = "NS to be registered with the DNS"
  value       = ({
    for key in var.projectid_list :
      key => {
        nameserver = module.dns-public-zone[key].name_servers
      }
    })
}
