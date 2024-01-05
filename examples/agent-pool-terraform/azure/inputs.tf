variable "azure_resource_group" {
  type        = string
  default     = ""
  description = "Azure location in which to deploy agents"
}

variable "agent_count" {
  type        = number
  description = "Number of agents to deploy"
}

// The teleport module performs validation and sets defaults.
variable "agent_roles" {
  type = list(string)
}

variable "insecure_direct_access" {
  type        = bool
  default     = false
  description = "Whether to enable direct access to agent instances. Only enable this in low-security demo environments."
}

variable "proxy_service_address" {
  type        = string
  description = "Host and HTTPS port of the Teleport Proxy Service"
}

variable "public_key_path" {
  type        = string
  description = "Path to a valid RSA public key with at least 2048 bits. The key is only used to pass validation in Azure, and is deleted from VMs created by this module."
  default     = ""
}

variable "region" {
  type        = string
  description = "Location in which to deploy agents (Azure location, AWS or GCP region)"
}

variable "subnet_id" {
  type        = string
  description = "Cloud provider subnet for deploying Teleport agents (subnet ID if using AWS or Azure, name or self link if using GCP)"
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
