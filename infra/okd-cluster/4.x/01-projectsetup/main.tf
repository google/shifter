/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
    "container.googleapis.com",
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
      "user:parasmamgain@google.com",                               ## Replace with useraccount who can impoersonate the SA
      "serviceAccount:1002225287836@cloudbuild.gserviceaccount.com" ## Replce the cloudbuild SA which can impersonate the SA
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
resource "google_cloudbuild_trigger" "createresource-trigger" {
  for_each      = toset(var.projectid_list)
  project       = each.key
  name          = "CreateOCPClusterShifterTrigger"
  description   = "This trigger initiates the OCP cluster resource deployment in GCP."
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
    _OKD_VERSION       = "4.10"
    _BILLING_ACCOUNT_ID = "01541A-27C980-D4B4C9"
    _PARENT             = "folders/808116942407"
    _DOMAIN             = "pm-gcp.com"
  }
  build {
    timeout       = "12000s"
    step {
      name       = local.ubuntu_builder
      entrypoint = "bash"
      id         = "create-okd-cluster"
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
          wget -q https://releases.hashicorp.com/terraform/$_TERRAFORM_VERSION/terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          unzip terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          mv terraform /usr/local/bin/ &&
          terraform version &&
          echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
          apt-get install -y apt-transport-https ca-certificates gnupg &&
          curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
          apt-get update && apt-get install -y google-cloud-sdk &&
          gcloud version &&
          cd infra/okd-cluster/4.x &&
          ./install.sh $_PROJECT_NAME $_CLUSTER_NAME $_OKD_VERSION  $_BILLING_ACCOUNT_ID $_PARENT $_DOMAIN &&
          cp -r /workspace/infra/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/* /persistent_volume/
        EOT
      ]
    }
    step {
      name       = local.ubuntu_builder
      id         = "deploy-shifter-binary"
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
            wget -q https://github.com/google/shifter/releases/download/$_SHIFTER_VERSION/shifter_linux_amd64 &&
            chmod +x shifter_linux_amd64 &&
            mv shifter_linux_amd64 /usr/local/bin/ &&
            mv /usr/local/bin/shifter_linux_amd64 /usr/local/bin/shifter
            shifter version &&
            echo "******************************************"
            echo "* Setting up shifter"
            echo "******************************************"
            source /persistent_volume/cluster_credentials.env &&
            mkdir -p /persistent_volume/shifter/output &&
            shifter cluster -e $$_CLUSTER_API_ENDPOINT_ -t $$_TOKEN_ convert --namespace default -o yaml /persistent_volume/shifter/output
        EOT
      ]
    }
    step {
      name       = "hashicorp/terraform:1.3.6"
      id         = "create-gke-cluster"
      entrypoint = "sh"
      args = [
        "-c",
        <<-EOT
            echo "******************************************"
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke init &&
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke plan -var="project_id=$_PROJECT_NAME" &&
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke apply -var="project_id=$_PROJECT_NAME" -auto-approve
        EOT
      ]
    }
    step {
      name       = local.ubuntu_builder
      id         = "deploy-gke-workload-from-shifter"
      entrypoint = "bash"
      volumes {
        name = "myvolume"
        path = "/persistent_volume"
      }
      args = [
        "-c",
        <<-EOT
          echo "******************************************"
          echo "* Installing gcloud"
          echo "******************************************"
          apt-get install -y unzip wget git curl &&
          echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
          apt-get install -y apt-transport-https ca-certificates gnupg &&
          curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
          apt-get update && apt-get install -y google-cloud-sdk &&
          gcloud version &&
          sudo apt-get install -y google-cloud-sdk-gke-gcloud-auth-plugin &&
          gke-gcloud-auth-plugin --version &&
          gcloud container clusters get-credentials gke-for-okd-workloads &&
          kubectl get pods --all-namespaces &&
          kubectl apply -f /persistent_volume/shifter/output --recursive &&
          kubectl get pods
        EOT
      ]
    }
    // add step to run the shifter public image
    // run the shifter against the cluster created above
    # artifacts {
    #   objects {
    #     location = "${module.gcs-automation[each.key].url}/builds/plan-file/$BRANCH_NAME/"
    #     paths = ["/workspace/infra/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/*",
    #     ]
    #   }
    # }
  }
  depends_on = [
    module.repository-shifter
  ]
  approval_config {
    approval_required = true
  }
}

/******************************************************************
      Delete OKD Cluster - Cloud Build Trigger to run the terraform code
*******************************************************************/
resource "google_cloudbuild_trigger" "deletecluster-trigger" {
  for_each      = toset(var.projectid_list)
  project       = each.key
  name          = "DeleteOkdCluster"
  description   = "This trigger initiates the deletion process of OCP cluster deployed in GCP."
  ignored_files = ["*"]
  included_files = [""]
  trigger_template {
    project_id  = each.key
    branch_name = "v0.3.1"
    repo_name   = "csr-shifter"
  }
  substitutions = {
    _TERRAFORM_VERSION = "1.1.5"
    _PROJECT_NAME       = "pm-singleproject-20"
    _CLUSTER_NAME       = "okd42"
    _OKD_VERSION        = "4.10"
    _BILLING_ACCOUNT_ID = "01541A-27C980-D4B4C9"
    _PARENT             = "folders/808116942407"
    _DOMAIN             = "pm-gcp.com"
  }
  build {
    timeout       = "12000s"
    step {
      name       = "hashicorp/terraform:1.3.6"
      id         = "delete-gke-cluster"
      entrypoint = "sh"
      args = [
        "-c",
        <<-EOT
            echo "******************************************"
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke init &&
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke plan -var="project_id=$_PROJECT_NAME" &&
            terraform -chdir=/workspace/infra/okd-cluster/4.x/03-gke destroy -var="project_id=$_PROJECT_NAME" -auto-approve
        EOT
      ]
    }
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
          wget -q https://releases.hashicorp.com/terraform/$_TERRAFORM_VERSION/terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          unzip terraform_$_TERRAFORM_VERSION\_linux_amd64.zip &&
          mv terraform /usr/local/bin/ &&
          terraform version &&
          echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list &&
          apt-get install -y apt-transport-https ca-certificates gnupg &&
          curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
          apt-get update && apt-get install -y google-cloud-sdk &&
          gcloud version &&
          cd infra/okd-cluster/4.x &&
          ./destroy.sh $_PROJECT_NAME $_CLUSTER_NAME $_OKD_VERSION
        EOT
      ]
    }
    # artifacts {
    #   objects {
    #     location = "${module.gcs-automation[each.key].url}/builds/plan-file/$BRANCH_NAME/"
    #     paths = ["/workspace/infra/okd-cluster/4.x/install-config/$_PROJECT_NAME/$_CLUSTER_NAME/.*",
    #     ]
    #   }
    # }
  }
  depends_on = [
    module.repository-shifter
  ]
  approval_config {
    approval_required = true
  }
}

module "gcs-automation" {
  for_each   = toset(var.projectid_list)
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/gcs?ref=v16.0.0"
  project_id = each.key
  name       = "shifter-tfstate"
  versioning = true
}
