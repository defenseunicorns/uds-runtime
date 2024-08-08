provider "aws" {
  region = var.region
}

resource "random_id" "unique_id" {
  byte_length = 4
}

data "aws_caller_identity" "current" {}

data "aws_ami" "latest_runtime_ephemeral_ami" {
  most_recent = true

  filter {
    name   = "name"
    values = ["runtime-ephemeral-*"]
  }

  owners = ["${data.aws_caller_identity.current.account_id}"]
}

locals {
  suffix = random_id.unique_id.hex
  tags = tomap({
    "Name"         = "runtime-ephemeral-${local.suffix}"
    "ManagedBy"    = "Terraform"
    "CreationDate" = time_static.creation_time.rfc3339
  })
}

resource "time_static" "creation_time" {}

resource "aws_instance" "ec2_instance" {
  ami           = data.aws_ami.latest_runtime_ephemeral_ami.image_id
  instance_type = "m5.2xlarge"
  tags          = local.tags

  vpc_security_group_ids = [aws_security_group.security_group.id]
  user_data              = file("setup.sh")

  root_block_device {
    volume_size           = 32
    volume_type           = "gp2"
    delete_on_termination = true
  }
}

resource "aws_security_group" "security_group" {
  name        = "runtime-ephemeral-sg-${random_id.unique_id.hex}"
  description = "kube-api access from anywhere"

  ingress {
    from_port   = 6550
    to_port     = 6550
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # ingress {
  #   from_port   = 6443
  #   to_port     = 6443
  #   protocol    = "tcp"
  #   cidr_blocks = ["0.0.0.0/0"]
  # }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.tags
}
