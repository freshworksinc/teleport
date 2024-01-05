// The root module performs validation and sets defaults.
variable "agent_roles" {
  type = list(string)
}

variable "proxy_service_address" {
  type        = string
  description = "Host and HTTPS port of the Teleport Proxy Service"
}

variable "teleport_edition" {
  type        = string
  default     = "oss"
  description = "Edition of your Teleport cluster. Can be: oss, enterprise, team, or cloud."
  validation {
    condition     = contains(["oss", "enterprise", "team", "cloud"], var.teleport_edition)
    error_message = "teleport_edition must be one of: oss, enterprise, team, cloud."
  }
}

variable "teleport_plugin_version" {
  type        = string
  description = "version of the Teleport Terraform provider to use"
}

variable "teleport_version" {
  type        = string
  description = "Version of Teleport to install on each agent"
}

