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
  iam = {
    "roles/iam.serviceAccountTokenCreator" = compact([
      "user:parasmamgain@google.com",
      "serviceAccount:1002225287836@cloudbuild.gserviceaccount.com"
    ])
  }

  iam_project_roles = {
    "${each.key}" = [
      "roles/owner",
      # "roles/dns.admin",
      # "roles/storage.admin",
      # "roles/container.admin",
      # "roles/compute.instanceAdmin",
      # "roles/compute.instanceAdmin.v1",
      # "roles/compute.networkAdmin",
      # "roles/compute.securityAdmin",
      # "roles/iam.serviceAccountUser",
      # "roles/iam.serviceAccountAdmin",
      # "roles/iam.serviceAccountKeyAdmin",
      # "roles/servicemanagement.quotaViewer",
      # "roles/resourcemanager.projectIamAdmin",
      "roles/secretmanager.secretAccessor",
      "roles/secretmanager.viewer",
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


#Creating Code repository and CLoudbuild pipeline
locals {
  ubuntu_builder    = "gcr.io/cloud-marketplace-containers/google/debian11"
  terraform_builder = "hashicorp/terraform:1.0.10"
  shifter_builder   = "images.shifter.cloud/shifter:latest"
}
module "repository-shifter" {
  for_each   = toset(var.projectid_list)
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/source-repository?ref=v16.0.0"
  project_id = each.key
  name       = "csr-shifter"
  iam = {
    //"roles/source.writer" = [module.sa-csr-writer.iam_email]
  }
}

/******************************************************************
      Create OKD Cluster - Cloud Build Trigger to run the terraform code
*******************************************************************/
resource "google_cloudbuild_trigger" "sharedresource-trigger" {
  for_each      = toset(var.projectid_list)
  project       = each.key
  name          = "ShifterTrigger"
  description   = "This trigger initiates the GCP resource deployment."
  ignored_files = ["*"]
  included_files = [""]
  trigger_template {
    project_id  = each.key
    branch_name = "v0.3.1"
    repo_name   = "csr-shifter"
  }
  substitutions = {
    _TERRAFORM_VERSION = "1.1.5"
    _SHIFTER_VERSION   = "v0.3.0"
    _PROJECT_NAME      = "pm-singleproject-20"
    _CLUSTER_NAME      = "okd42"
  }
  build {
    timeout       = "4200s"
    step {
      name       = local.ubuntu_builder
      entrypoint = "bash"
      volumes {
        name = "myvolume"
        path = "/persistent_volume"
      }
      args = [
        "-c",
        <<-EOT
          echo "******************************************"
          echo "* Installing Terraform,gcloud"
          echo "******************************************"
          apt-get install -y unzip wget git curl &&
          wget https://releases.hashicorp.com/terraform/$_TERRAFORM_VERSION/terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          unzip terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          mv terraform /usr/local/bin/ &&
          terraform version &&
          echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
          apt-get install -y apt-transport-https ca-certificates gnupg &&
          curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
          apt-get update && apt-get install -y google-cloud-sdk &&
          gcloud version &&
          cd okd-cluster/4.x &&
          ./install.sh &&
          cp -r /workspace/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/* /persistent_volume/
        EOT
      ]
    }
    step {
      name       = local.ubuntu_builder
      env        = ["TOKEN = ","CLUSTER_API_ENDPOINT= "]
      entrypoint = "bash"
      volumes {
        name = "myvolume"
        path = "/persistent_volume"
      }
      args = [
        "-c",
        <<-EOT
            echo "******************************************"
            echo "* Installing Shifter"
            echo "******************************************"
            apt-get install -y unzip wget git curl &&
            wget https://github.com/google/shifter/releases/download/$_SHIFTER_VERSION/shifter_linux_amd64 &&
            chmod +x shifter_linux_amd64 &&
            mv shifter_linux_amd64 /usr/local/bin/ &&
            mv /usr/local/bin/shifter_linux_amd64 /usr/local/bin/shifter
            shifter version &&
            ls -R /persistent_volume
            echo "******************************************"
            echo "* Setting up Cluster"
            echo "******************************************"
            source /persistent_volume/cluster_credentials.env
            shifter cluster -e $CLUSTER_API_ENDPOINT -t $TOKEN list
        EOT
      ]
    }
    // add step to run the shifter public image
    // run the shifter against the cluster created above
    # artifacts {
    #   objects {
    #     location = "${module.gcs-automation[each.key].url}/builds/plan-file/$BRANCH_NAME/"
    #     paths = ["/workspace/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/*",
    #     ]
    #   }
    # }
  }
  depends_on = [
    module.repository-shifter
  ]
}

/******************************************************************
      Delete OKD Cluster - Cloud Build Trigger to run the terraform code
*******************************************************************/
resource "google_cloudbuild_trigger" "deletecluster-trigger" {
  for_each      = toset(var.projectid_list)
  project       = each.key
  name          = "DeleteOkdCluster"
  description   = "This trigger initiates the GCP resource deployment."
  ignored_files = ["*"]
  included_files = [""]
  trigger_template {
    project_id  = each.key
    branch_name = "v0.3.1"
    repo_name   = "csr-shifter"
  }
  substitutions = {
    _TERRAFORM_VERSION = "1.1.5"
    _PROJECT_NAME      = "pm-singleproject-20"
    _CLUSTER_NAME      = "okd42"
  }
  build {
    timeout       = "3600s"
    step {
      name       = local.ubuntu_builder
      entrypoint = "bash"
      args = [
        "-c",
        <<-EOT
          echo "******************************************"
          echo "* Installing Terraform,gcloud"
          echo "******************************************"
          apt-get install -y unzip wget git curl &&
          wget https://releases.hashicorp.com/terraform/$_TERRAFORM_VERSION/terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          unzip terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          mv terraform /usr/local/bin/ &&
          terraform version &&
          echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
          apt-get install -y apt-transport-https ca-certificates gnupg &&
          curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
          apt-get update && apt-get install -y google-cloud-sdk &&
          gcloud version &&
          cd okd-cluster/4.x &&
          ./destroy.sh
        EOT
      ]
    }
    # artifacts {
    #   objects {
    #     location = "${module.gcs-automation[each.key].url}/builds/plan-file/$BRANCH_NAME/"
    #     paths = ["/workspace/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/.*",
    #     ]
    #   }
    # }
  }
  depends_on = [
    module.repository-shifter
  ]
}

module "gcs-automation" {
  for_each   = toset(var.projectid_list)
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/gcs?ref=v16.0.0"
  project_id = each.key
  name       = "shifter-tfstate"
  versioning = true
}
