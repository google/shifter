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
