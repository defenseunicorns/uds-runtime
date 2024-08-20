variable "region" {
  description = "The AWS region to deploy the EC2 instance"
  type        = string
  default     = "us-west-2"
}

variable "permissions_boundary_name" {
  description = "The name of the permissions boundary to attach to the IAM role"
  type        = string
}

variable "permissions_boundary_arn" {
  description = "The ARN of the permissions boundary to attach to the IAM role"
  type        = string
}

variable "enable_ssh" {
  description = "Enable SSH access to the EC2 instance"
  type        = bool
  default     = false
}

variable "ssh_ip" {
  description = "The IP address to allow SSH access from"
  type        = string
  default     = ""
}
