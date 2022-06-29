resource "google_compute_firewall" "openshift" {
  project     = module.project.project_id
  name        = "${var.prefix}-cluster-fw"
  network     = module.network.vpc_network.0.name
  description = "Allow Openshift comms"

  allow {
    protocol = "all"
  }

  source_tags = ["openshift"]
  target_tags = ["openshift"]
}

resource "google_compute_address" "os-master-addr" {
  project = module.project.project_id
  name    = "${var.prefix}-cluster-master-addr"
  region  = var.region
}

resource "google_compute_http_health_check" "os-master-hc" {
  project             = module.project.project_id
  name                = "${var.prefix}-cluster-master-hc"
  check_interval_sec  = 10
  healthy_threshold   = 10
  unhealthy_threshold = 10
  timeout_sec         = 5
  request_path        = "/"
  port                = "8443"
}

resource "google_compute_firewall" "os-master-hc-fw" {
  project     = module.project.project_id
  name        = "${var.prefix}-cluster-hc-fw"
  network     = module.network.vpc_network.0.name
  description = "Allow HelathCheck comms"

  allow {
    protocol = "tcp"
    ports    = ["8443"]
  }

  source_ranges = ["35.191.0.0/16", "130.211.0.0/22"]
  target_tags   = ["os-cp"]
}

resource "google_compute_target_pool" "os-master-tp" {
  project       = module.project.project_id
  name          = "${var.prefix}-cluster-master-tp"
  instances     = [google_compute_instance_from_template.os-control-plane.0.self_link]
  health_checks = [google_compute_http_health_check.os-master-hc.name]
  region        = var.region
}

resource "google_compute_forwarding_rule" "os-master-https" {
  project    = module.project.project_id
  name       = "${var.prefix}-cluster-https-fwd-rule"
  target     = google_compute_target_pool.os-master-tp.self_link
  ip_address = google_compute_address.os-master-addr.address
  port_range = "8443"
  region     = var.region
}

