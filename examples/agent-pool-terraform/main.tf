module "teleport" {
  source                = "./teleport"
  proxy_service_address = var.proxy_service_address
  teleport_edition      = var.teleport_edition
  teleport_version      = var.teleport_version
}

module "aws" {
  count                   = var.cloud == "aws" ? 1 : 0
  agent_count             = var.agent_count
  agent_roles             = var.agent_roles
  proxy_service_address   = var.proxy_service_address
  region                  = var.region
  source                  = "./aws"
  subnet_id               = var.subnet_id
  teleport_edition        = var.teleport_edition
  teleport_version        = var.teleport_version
  userdata                = module.teleport.userdata
}

module "gcp" {
  count                   = var.cloud == "gcp" ? 1 : 0
  agent_count             = var.agent_count
  agent_roles             = var.agent_roles
  gcp_zone                = var.gcp_zone
  google_project          = var.google_project
  proxy_service_address   = var.proxy_service_address
  region                  = var.region
  source                  = "./gcp"
  subnet_id               = var.subnet_id
  teleport_edition        = var.teleport_edition
  teleport_version        = var.teleport_version
  userdata                = module.teleport.userdata
}

module "azure" {
  count                   = var.cloud == "azure" ? 1 : 0
  agent_count             = var.agent_count
  agent_roles             = var.agent_roles
  azure_resource_group    = var.azure_resource_group
  proxy_service_address   = var.proxy_service_address
  public_key_path         = var.public_key_path
  region                  = var.region
  source                  = "./azure"
  subnet_id               = var.subnet_id
  teleport_edition        = var.teleport_edition
  teleport_version        = var.teleport_version
  userdata                = module.teleport.userdata
}
