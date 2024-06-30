variable "name" {
  type = string
  default = "default value"
}

variable "my_map" {
  type = map(any)
}

terraform {
  required_providers {
    local = {
      source = "hashicorp/local"
      version = "2.4.0"
    }
  }
}

module "write_local_file" {
  source = "./modules"

  name = var.name
  my_map = var.my_map
}