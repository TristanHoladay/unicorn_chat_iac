terraform {
  required_version = ">= 1.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.62.0"
    }
  }
}

provider "aws" {
  region                   = var.region
  profile                  = var.profile
  shared_config_files      = [var.aws_config_path]
  shared_credentials_files = [var.aws_creds_path]
}
