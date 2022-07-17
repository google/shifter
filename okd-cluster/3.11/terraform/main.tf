module "network" {
  source         = "./modules/gcp_tf_network"
  project_id     = module.project.project_id
  prefix         = var.prefix
  name           = "1"
  default_region = var.region

  subnets = [
    {
      name       = "eu1"
      region     = var.region
      cidr_range = "10.0.0.0/22"
    },
  ]
}

