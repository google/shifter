prefix       = "dev"
region       = "europe-west1"
master_count = 1
node_count   = 2

project_name              = "okd-tf"
# Need to be passed only when create_project variable is set to false else leave it as empty variable
project_id                = ""
org_id                    = "615056687435"
folder_id                 = ""
environment               = "development"
billing_code              = "1234"
billing_account           = "0090FE-ED3D81-AF8E3B"
application_name          = ""
primary_contact           = "avinashjha@google.com"
activate_apis = [
  "compute.googleapis.com",
  "cloudbilling.googleapis.com",
  "dns.googleapis.com",
  "servicenetworking.googleapis.com"
]
master_subdomain = "okd.avinashj.in"
public_subdomain = "console.okd.avinashj.in"
dns_master_subdomain = "okd.avinashj.in."
ssh_user = "avinashjha"