variable "region" {
  description = "The AWS region to deploy the EC2 instance"
  type        = string
  default    = "us-east-1"
}

variable "enable_ssh" {
  description = "Enable SSH access to the EC2 instance"
  type        = bool
  default     = false
}

variable "ssh_ip" {
  description = "The IP address to allow SSH access from"
  type        = string
  default    = ""
}
