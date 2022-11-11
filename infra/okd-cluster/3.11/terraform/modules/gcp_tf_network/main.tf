resource "google_compute_network" "vpc" {
  name                    = "${var.prefix}-${var.name}-vpc"
  project                 = var.project_id
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  count                    = length(var.subnets)
  project                  = var.project_id
  network                  = google_compute_network.vpc.self_link
  name                     = "${var.prefix}-${var.name}-${var.subnets[count.index]["name"]}"
  region                   = var.subnets[count.index]["region"]
  ip_cidr_range            = var.subnets[count.index]["cidr_range"]
  private_ip_google_access = true
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "${var.prefix}-${var.name}-private-ip-alloc"
  project       = var.project_id
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.vpc.self_link
}

resource "google_service_networking_connection" "svc-network-peering" {
  network                 = google_compute_network.vpc.self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}

/* ------ Create Cloud NAT Router ------ */
resource "google_compute_router" "nat_router" {
  count = length(var.subnets)

  name    = "${var.prefix}-${var.name}-nat-rtr-${count.index}"
  project = var.project_id
  region  = var.subnets[count.index]["region"]
  network = google_compute_network.vpc.self_link

  bgp {
    asn = 64514
  }
}

resource "google_compute_address" "nat_gw_address" {
  project = var.project_id
  count   = length(var.subnets)
  name    = "${var.prefix}-${var.name}-nat-ext-addr-${count.index}"
  region  = element(google_compute_subnetwork.subnet.*.region, count.index)
}

resource "google_compute_router_nat" "nat_gateway" {
  count   = length(var.subnets)
  project = var.project_id
  name    = "${var.prefix}-${var.name}-nat-gw-${count.index}"

  router = element(google_compute_router.nat_router.*.name, count.index)
  region = element(google_compute_subnetwork.subnet.*.region, count.index)

  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [element(google_compute_address.nat_gw_address.*.self_link, count.index)]
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"

  subnetwork {
    name                    = element(google_compute_subnetwork.subnet.*.self_link, count.index)
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }

  log_config {
    filter = "TRANSLATIONS_ONLY"
    enable = true
  }
}

resource "google_compute_firewall" "deny-all" {
  project = var.project_id
  name    = "${var.prefix}-${var.name}-deny-all"
  network = google_compute_network.vpc.name

  priority = 65535

  deny {
    protocol = "all"
  }

  source_ranges = ["0.0.0.0/0"]

}

resource "google_compute_firewall" "ingress-allow-iap" {
  project = var.project_id
  name    = "${var.prefix}-${var.name}-ingress-allow-iap"
  network = google_compute_network.vpc.name

  priority = 200

  allow {
    protocol = "tcp"
  }

  source_ranges = ["35.235.240.0/20"]
}
