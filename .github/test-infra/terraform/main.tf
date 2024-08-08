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
  key_name = var.enable_ssh ? aws_key_pair.ssh[0].key_name: null
  tags          = local.tags

  vpc_security_group_ids = [aws_security_group.security_group.id]
  user_data              = file("setup.sh")

  root_block_device {
    volume_size           = 32
    volume_type           = "gp2"
    delete_on_termination = true
  }
}

#
# SSH Config for Testing
#
resource "tls_private_key" "ssh" {
  count = var.enable_ssh ? 1 : 0

  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "local_file" "ssh_pem" {
  count = var.enable_ssh ? 1 : 0

  filename        = "runtime-dev.pem"
  content         = tls_private_key.ssh[0].private_key_pem
  file_permission = "0400"
}

resource "local_file" "ssh_pub" {
  count = var.enable_ssh ? 1 : 0

  filename        = "runtime-dev.pub"
  content         = tls_private_key.ssh[0].public_key_openssh
  file_permission = "0644"
}

resource "aws_key_pair" "ssh" {
  count = var.enable_ssh ? 1 : 0

  key_name   = "runtime-dev-key"
  # public_key = tls_private_key.ssh[0].public_key_openssh
   public_key = local_file.ssh_pub[0].content
}


resource "aws_security_group" "security_group" {
  name        = "runtime-ephemeral-sg-${random_id.unique_id.hex}"

   ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["${var.ssh_ip}/32"]
  }

  ingress {
    from_port   = 6550
    to_port     = 6550
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = local.tags
}
