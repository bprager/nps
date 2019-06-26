# variables.tf
variable "workspace_to_environment_map" {
  type = "map"
  default = {
    dev  = "dev"
    prod = "prod"
  }
}
