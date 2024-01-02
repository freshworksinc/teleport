module "teleport" {
  source = "../teleport"
  proxy_service_address = var.proxy_service_address
  teleport_edition = var.teleport_edition
  teleport_version = var.teleport_version
}

resource "google_compute_instance" "teleport_agent" {
  count = var.agent_count
  name  = "teleport-agent-${count.index}"
  zone  = var.gcp_zone

  boot_disk {
    initialize_params {
      image = "family/ubuntu-2204-lts"
    }
  }

  network_interface {
    subnetwork = var.subnet_id
  }

  machine_type = "e2-standard-2"

  metadata_startup_script = module.teleport.userdata 
}
