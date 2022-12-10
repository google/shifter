


module "vpc" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric/modules/net-vpc"
  project_id = var.project_id
  name       = var.vpc_network_name
  vpc_create = var.vpc_create
  subnets = [
    {
      ip_cidr_range = var.vpc_subnetwork_cidr
      name          = var.vpc_subnetwork_name
      region        = var.region
      secondary_ip_ranges = {
        pods     = var.gke_pods_secondary_cidr
        services = var.gke_services_secondary_cidr
      }
    }
  ]
}

module "gke" {
  source                     = "terraform-google-modules/kubernetes-engine/google"
  project_id                 = var.project_id
  name                       = var.gke_cluster_name
  region                     = var.region
  zones                      = ["${var.region}-a", "${var.region}-f", "${var.region}-b"]
  network                    = var.vpc_network_name
  subnetwork                 = var.vpc_subnetwork_name
  ip_range_pods              = "pods"
  ip_range_services          = "services"
  http_load_balancing        = false
  network_policy             = false
  horizontal_pod_autoscaling = true
  filestore_csi_driver       = false

  node_pools = [
    {
      name            = "default-node-pool"
      machine_type    = "e2-medium"
      node_locations  = "${var.region}-a,${var.region}-f"
      min_count       = 1
      max_count       = 12
      local_ssd_count = 0
      spot            = false
      disk_size_gb    = 40
      disk_type       = "pd-standard"
      image_type      = "COS_CONTAINERD"
      enable_gcfs     = false
      enable_gvnic    = false
      auto_repair     = true
      auto_upgrade    = true
      #service_account           = "project-service-account@<PROJECT ID>.iam.gserviceaccount.com"
      preemptible        = false
      initial_node_count = 3
    },
  ]

  auto_scaling = {
    all ={
      location_policy = "BALANCED"
    }
  }

  node_pools_oauth_scopes = {
    all = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }

  node_pools_labels = {
    all = {}

    default-node-pool = {
      default-node-pool = true
    }
  }

  node_pools_metadata = {
    all = {}

    default-node-pool = {
      node-pool-metadata-custom-value = "my-node-pool"
    }
  }

  node_pools_taints = {
    all = []

    default-node-pool = [
      {
        key    = "default-node-pool"
        value  = true
        effect = "PREFER_NO_SCHEDULE"
      },
    ]
  }

  node_pools_tags = {
    all = []

    default-node-pool = [
      "default-node-pool",
    ]
  }

  depends_on = [
    module.vpc
  ]
}
