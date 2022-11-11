
#Update the service accounts
provider "google" {
  impersonate_service_account = "okd-sa@PROJECT-ID.iam.gserviceaccount.com"
}
provider "google-beta" {
  impersonate_service_account = "okd-sa@PROJECT-ID.iam.gserviceaccount.com"
}
terraform {
  backend "gcs" {
    bucket                      = "shifter-tfstate"
    prefix                      = "shifter/4.x"
    impersonate_service_account = "okd-sa@PROJECT-ID.iam.gserviceaccount.com"
  }
}
