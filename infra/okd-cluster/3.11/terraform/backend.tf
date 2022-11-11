terraform {
  backend "gcs" {
    bucket = "okd-tf01"
    prefix = "terraform/okd/demo"
  }
}
