variable "region" {
  description = "Region of S3 bucket"
  type        = string
  default     = "us-east-1"
}

variable "profile" {
  description = "AWS profile to use for Terraform"
  default     = ""
}

variable "aws_config_path" {
  description = "Path to local AWS config for authentication."
  default     = ""
}

variable "aws_creds_path" {
  description = "Path to local AWS credentials for authentication."
  default     = ""
}
