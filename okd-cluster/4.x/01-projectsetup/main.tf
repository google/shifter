//Removes org policies at project level and enabling required API's
module "project" {
  for_each        = toset(var.projectid_list)
  source          = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/project?ref=v16.0.0"
  billing_account = var.billing_account_id
  name            = each.key
  parent          = var.parent
  prefix          = null
  project_create  = var.project_create
  services = [
    "compute.googleapis.com",
    "cloudapis.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "dns.googleapis.com",
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
      values              = ["under:projects/${each.key}"]
    },
  }
  service_config = {
    disable_on_destroy         = false
    disable_dependent_services = false
  }
}

//Creates dns public zone in the respective projects, to be used for deploying the cluster
module "dns-public-zone" {
  for_each   = toset(var.projectid_list)
  source     = "terraform-google-modules/cloud-dns/google"
  version    = "3.0.0"
  project_id = each.key
  type       = "public"
  name       = "${each.key}-zone-name"
  domain     = "${each.key}.${var.domain}"
  recordsets = [
  ]
  depends_on = [
    module.project
  ]
}

//Creates service account with owner permission to be used by the openshift-installer for deployment
module "okd-sa" {
  for_each     = toset(var.projectid_list)
  source       = "github.com/terraform-google-modules/cloud-foundation-fabric//modules/iam-service-account?ref=v15.0.0"
  project_id   = each.key
  name         = "okd-sa"
  generate_key = false
  # non-authoritative roles granted *to* the service accounts on other resources
  iam_project_roles = {
    "${each.key}" = [
      "roles/owner",
      "roles/dns.admin",
      "roles/compute.viewer",
      "roles/storage.admin",
      "roles/compute.instanceAdmin",
      "roles/compute.networkAdmin",
      "roles/compute.securityAdmin",
      "roles/iam.serviceAccountAdmin",
      "roles/iam.serviceAccountUser",
      "roles/iam.serviceAccountKeyAdmin",
      "roles/servicemanagement.quotaViewer",
      "roles/resourcemanager.projectIamAdmin",
    ]
  }
  depends_on = [
    module.dns-public-zone,
    module.project
  ]
}

//Creates yaml file for deploying the openshift cluster
data "template_file" "install_config_yaml" {
  for_each = toset(var.projectid_list)
  template = file("template/install-config-template.yaml")
  vars = {
    dns-public-zone    = "${each.key}.${var.domain}"
    metadata-name      = "${each.key}"
    cluster-name       = "${var.cluster_name}"
    region             = "${var.region}"
    pull-secret-redhat = "${var.redhat_pull_secret}"
    ssh-key            = var.ssh_key_path == "" ? "" : data.local_file.ssh_pub.0.content
  }
}
resource "local_file" "init_script" {
  for_each = toset(var.projectid_list)
  content  = data.template_file.install_config_yaml[each.key].rendered
  filename = "../install-config/${each.key}/${var.cluster_name}/install-config.yaml"
}

data "local_file" "ssh_pub" {
  count    = var.ssh_key_path == "" ? 0 : 1
  filename = var.ssh_key_path
}
