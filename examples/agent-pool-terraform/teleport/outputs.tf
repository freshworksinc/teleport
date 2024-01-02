locals {
  userdata = templatefile("./userdata", {
    token                 = teleport_provision_token.agent[count.index].metadata.name
    proxy_service_address = var.proxy_service_address
    teleport_edition      = var.teleport_edition
    teleport_version      = var.teleport_version
  }
}

output "userdata" {
  value = local.userdata
  description = "User data script to run on agent instances."
}
