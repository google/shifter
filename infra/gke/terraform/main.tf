


module "project" {
  source          = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/project"
  billing_account = var.billing_account_id
  name            = var.project_id
  parent          = var.parent
  prefix          = null
  project_create  = var.project_create
  services = [
    "compute.googleapis.com",
    "cloudapis.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "container.googleapis.com",
    "iamcredentials.googleapis.com",
    "iam.googleapis.com",
    "servicemanagement.googleapis.com",
    "serviceusage.googleapis.com",
    "storage-api.googleapis.com",
    "storage-component.googleapis.com",
  ]
  policy_boolean = {
    "constraints/iam.disableServiceAccountKeyCreation" = false
    "constraints/compute.skipDefaultNetworkCreation"   = true
  }
  policy_list = {
    "constraints/compute.restrictLoadBalancerCreationForTypes" = {
      inherit_from_parent = null
      suggested_value     = null
      status              = true
      values              = []
    },
    "constraints/compute.vmExternalIpAccess" = {
      inherit_from_parent = null
      suggested_value     = null
      status              = true
      values              = []
    },
    "constraints/compute.restrictCloudNATUsage" = {
      inherit_from_parent = null
      suggested_value     = null
      status              = true
      values              = ["under:projects/${var.project_id}"]
    },
  }
  service_config = {
    disable_on_destroy         = false
    disable_dependent_services = false
  }
}

module "vpc" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric/modules/net-vpc"
  project_id = module.project.project_id
  name       = var.vpc_network_name
  vpc_create = var.vpc_create
  subnets = [
    {
      ip_cidr_range = var.vpc_subnetwork_cidr
      name          = var.vpc_subnetwork_name
      region        = var.region
      secondary_ip_range = {
        pods     = var.gke_pods_secondary_cidr
        services = var.gke_services_secondary_cidr
      }
    }
  ]
}

module "gke" {
  source                    = "github.com/GoogleCloudPlatform/cloud-foundation-fabric/modules/gke-cluster"
  project_id                = module.project.project_id
  name                      = var.gke_cluster_name
  location                  = var.gke_location
  network                   = module.vpc.self_link
  subnetwork                = module.vpc.subnet_self_links["${var.region}/${var.vpc_subnetwork_name}"]
  secondary_range_pods      = "pods"
  secondary_range_services  = "services"
  default_max_pods_per_node = 32
  labels = {
    environment = "development"
  }
}

module "gke-nodepool" {
  source                      = "github.com/GoogleCloudPlatform/cloud-foundation-fabric/modules/gke-nodepool"
  project_id                  = module.project.project_id
  cluster_name                = module.gke.name
  location                    = module.gke.location
  name                        = var.gke_nodepool_name
  node_count	                = var.gke_nodepool_count
  node_service_account_create = true
}
