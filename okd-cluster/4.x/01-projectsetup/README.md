<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_local"></a> [local](#provider\_local) | 2.2.3 |
| <a name="provider_template"></a> [template](#provider\_template) | 2.2.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_dns-public-zone"></a> [dns-public-zone](#module\_dns-public-zone) | terraform-google-modules/cloud-dns/google | 3.0.0 |
| <a name="module_okd-sa"></a> [okd-sa](#module\_okd-sa) | github.com/terraform-google-modules/cloud-foundation-fabric//modules/iam-service-account | v15.0.0 |
| <a name="module_project"></a> [project](#module\_project) | github.com/GoogleCloudPlatform/cloud-foundation-fabric//modules/project | v16.0.0 |

## Resources

| Name | Type |
|------|------|
| [local_file.init_script](https://registry.terraform.io/providers/hashicorp/local/latest/docs/resources/file) | resource |
| [local_file.ssh_pub](https://registry.terraform.io/providers/hashicorp/local/latest/docs/data-sources/file) | data source |
| [template_file.install_config_yaml](https://registry.terraform.io/providers/hashicorp/template/latest/docs/data-sources/file) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_billing_account_id"></a> [billing\_account\_id](#input\_billing\_account\_id) | Billing account id to be associated with the project. | `string` | n/a | yes |
| <a name="input_cluster_name"></a> [cluster\_name](#input\_cluster\_name) | Name of OKD Cluster(max length restricted to 10) | `string` | n/a | yes |
| <a name="input_disable_dependent_services"></a> [disable\_dependent\_services](#input\_disable\_dependent\_services) | Whether services that are enabled and which depend on this service should also be disabled when this service is destroyed. https://www.terraform.io/docs/providers/google/r/google_project_service.html#disable_dependent_services | `bool` | `false` | no |
| <a name="input_disable_services_on_destroy"></a> [disable\_services\_on\_destroy](#input\_disable\_services\_on\_destroy) | Whether project services will be disabled when the resources are destroyed. https://www.terraform.io/docs/providers/google/r/google_project_service.html#disable_on_destroy | `bool` | `false` | no |
| <a name="input_domain"></a> [domain](#input\_domain) | The domain name owned by the user which will then be used for creating a public facing cluster. | `string` | n/a | yes |
| <a name="input_enable_apis"></a> [enable\_apis](#input\_enable\_apis) | Whether to actually enable the APIs. If false, this module is a no-op. | `bool` | `true` | no |
| <a name="input_parent"></a> [parent](#input\_parent) | parent orgId under which the project exists(or will be created). | `string` | n/a | yes |
| <a name="input_project_create"></a> [project\_create](#input\_project\_create) | This flag will help to either create a new project or use a project that already exists. | `bool` | `false` | no |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | The GCP project you want to enable APIs and create your project. | `string` | `""` | no |
| <a name="input_projectid_list"></a> [projectid\_list](#input\_projectid\_list) | The GCP project you want to enable APIs and create your project. | `list(string)` | n/a | yes |
| <a name="input_redhat_pull_secret"></a> [redhat\_pull\_secret](#input\_redhat\_pull\_secret) | Redhat pull secret(default it uses a generic pull secret). | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | Region where the cluster and its resources will be created. | `string` | `"us-central1"` | no |
| <a name="input_ssh_key_path"></a> [ssh\_key\_path](#input\_ssh\_key\_path) | Path for the ssh key public certificate for the machine which can be used to troubleshoot the clusters. | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_nameservers"></a> [nameservers](#output\_nameservers) | NS to be registered with the DNS |
<!-- END_TF_DOCS -->