terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }

    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.0"
    }

    google = {
      source  = "hashicorp/google"
      version = "~> 5.5.0"
    }

    teleport = {
      source = "terraform.releases.teleport.dev/gravitational/teleport"
    }
  }
}
