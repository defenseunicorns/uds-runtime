packer {
  required_version = ">= 1.8.7"

  required_plugins {
    amazon = {
      version = ">= 1.1.6"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

locals {
  ami_name = "runtime-ephemeral-${formatdate("YYYYMMDDhhmm", timestamp())}"
}

source "amazon-ebs" "ubuntu" {
  ami_name        = local.ami_name
  ami_description = "For testing uds-runtime releases"
  instance_type   = "t3a.medium"
  region          = "us-east-1"
  ssh_username    = "ubuntu"
  # ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20230608
  source_ami = "ami-0c7217cdde317cfec"
}

build {
  name    = local.ami_name
  sources = ["source.amazon-ebs.ubuntu"]

  # wait for cloud-init to finish before running the install script
  provisioner "shell" {
    inline = [
     "/usr/bin/cloud-init status --wait",
     "echo set debconf to Noninteractive",
     "echo 'debconf debconf/frontend select Noninteractive' | sudo debconf-set-selections"
    ]
    timeout = "5m"
  }

  # install tools
  provisioner "shell" {
    script  = "./install-tools.sh"
    timeout = "15m"
  }
}
