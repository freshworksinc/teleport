required_providers {
  teleport = {
    source  = "terraform.releases.teleport.dev/gravitational/teleport"
    version = var.teleport_plugin_version
  }
}
