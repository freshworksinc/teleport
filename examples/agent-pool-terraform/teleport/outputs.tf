locals {
  userdata = templatefile("./userdata", {
    agent_config = templatefile("${path.root}/agent-config.yaml", {
      proxy_service_address = var.proxy_service_address
    })
    teleport_edition = var.teleport_edition
    teleport_version = var.teleport_version
    token            = teleport_provision_token.agent[count.index].metadata.name
  })
}

output "userdata" {
  value       = local.userdata
  description = "User data script to run on agent instances."
}
