
#Update the service accounts
provider "google" {
  impersonate_service_account = "okd-sa@pm-singleproject-20.iam.gserviceaccount.com"
}
provider "google-beta" {
  impersonate_service_account = "okd-sa@pm-singleproject-20.iam.gserviceaccount.com"
}
terraform {
  backend "gcs" {
    bucket                      = "shifter-tfstate"
    prefix                      = "shifter/4.x"
    impersonate_service_account = "okd-sa@pm-singleproject-20.iam.gserviceaccount.com"
  }
}

