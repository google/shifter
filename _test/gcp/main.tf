data "google_compute_zones" "europe-west1" {
  project = var.project_id
  region  = var.region
}

resource "random_shuffle" "az" {
  input        = data.google_compute_zones.europe-west1.names
  result_count = 1
}

resource "google_compute_project_metadata" "oslogin" {
  project = var.project_id
  metadata = {
    enable-oslogin = "TRUE"
  }
  lifecycle {
    ignore_changes = [
      metadata
    ]
  }
}

resource "google_project_service" "project_apis" {
  project = var.project_id
  count   = length(var.services)
  service = element(var.services, count.index)

  disable_dependent_services = true
  disable_on_destroy         = true
}

module "prod-network" {
  source         = "./modules/gcp_tf_network"
  project_id     = var.project_id
  prefix         = var.prefix
  name           = var.env
  default_region = var.region

  subnets = [
    {
      name       = "eu1"
      region     = "europe-west1"
      cidr_range = "10.76.36.0/22"
    },
  ]
  depends_on = [google_project_service.project_apis, ]
}

module "gke" {
  source     = "github.com/garybowers/gcp_tf_privategke"
  project_id = var.project_id
  name       = "${var.env}-a"

  region   = var.region
  zone     = random_shuffle.az.result[0]
  prefix   = var.prefix
  location = "europe-west1"

  master_ipv4_cidr_block = "192.168.1.0/28"

  vpc_network = module.prod-network.vpc_network.0.self_link
  subnet      = module.prod-network.subnetwork.0.self_link

  gke_min_version = "1.18"
  min_nodes       = 1
  max_nodes       = 10

  node_pools = [
    {
      name         = "np-a"
      machine_type = "e2-standard-2"
      min_count    = 1
      max_count    = 10

      image_type         = "COS_CONTAINERD"
      auto_repair        = true
      auto_upgrade       = true
      preemptible        = false
      initial_node_count = 1

      disk_type    = "pd-standard"
      disk_size_gb = 100
    }
  ]
  whitelist_ips = [
    {
      cidr_block   = "0.0.0.0/0"
      display_name = "all"
    },
  ]
}
