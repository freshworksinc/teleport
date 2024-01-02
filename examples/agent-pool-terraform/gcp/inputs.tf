variable "agent_count" {
  type        = number
  description = "Number of agents to deploy"
}

variable "cloud" {
  type        = string
  description = "Cloud provider: aws|gcp|azure"
  validation {
    condition     = var.cloud == "aws" || var.cloud == "gcp" || var.cloud == "azure"
    error_message = "The value of \"cloud\" must be \"aws\", \"gcp\", or \"azure\"."
  }
}

variable "google_project" {
  type        = string
  default     = ""
  description = "GCP project to associate agents with"
}

variable "gcp_zone" {
  type        = string
  default     = ""
  description = "GCP zone to associate agents with"
}

variable "proxy_service_address" {
  type        = string
  description = "Host and HTTPS port of the Teleport Proxy Service"
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

variable "teleport_version" {
  type        = string
  description = "Version of Teleport to install on each agent"
}

