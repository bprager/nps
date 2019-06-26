# main.tf
provider "kubernetes" {}

locals {
  environment = "${lookup(var.workspace_to_environment_map, terraform.workspace, "dev")}"
}
