module "project" {
  source                      = "terraform-google-modules/project-factory/google"
  version                     = "~> 10.1"
  random_project_id           = true
  org_id                      = var.org_id
  activate_apis               = distinct(concat(var.activate_apis, ["billingbudgets.googleapis.com"]))
  name                        = var.project_name
  billing_account             = var.billing_account
  folder_id                   = var.folder_id

  labels = {
    environment       = var.environment
    application_name  = var.application_name
    billing_code      = var.billing_code
    primary_contact   = element(split("@", var.primary_contact), 0)
    secondary_contact = element(split("@", var.secondary_contact), 0)
    business_code     = var.business_code
    vpc_type          = var.vpc_type
  }
}