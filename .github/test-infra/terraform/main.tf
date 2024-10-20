# Copyright 2024 Defense Unicorns
# SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

provider "aws" {
  region = var.region

  default_tags {
    tags  = {
    Name         = "runtime-ephemeral-${random_id.unique_id.hex}"
    ManagedBy   = "Terraform"
    CreationDate = time_static.creation_time.rfc3339
    nuke = "DO-NOT-DELETE"
    PermissionsBoundary = "${var.permissions_boundary_name}"
    }
  }
}

resource "random_id" "unique_id" {
  byte_length = 4
}

data "aws_partition" "current" {}

data "aws_caller_identity" "current" {}

data "aws_ami" "latest_runtime_ephemeral_ami" {
  most_recent = true

  filter {
    name   = "name"
    values = ["runtime-ephemeral-*"]
  }

  owners = ["${data.aws_caller_identity.current.account_id}"]
}

resource "time_static" "creation_time" {}

#
# EC2 INSTANCE
#
resource "aws_instance" "runtime" {
  ami                  = data.aws_ami.latest_runtime_ephemeral_ami.image_id
  instance_type        = "m5.2xlarge"
  iam_instance_profile = aws_iam_instance_profile.runtime_profile.name
  key_name             = var.enable_ssh ? aws_key_pair.ssh[0].key_name : null

  vpc_security_group_ids = [aws_security_group.security_group.id]
  user_data              = file("setup.sh")
  root_block_device {
    volume_size           = 32
    volume_type           = "gp2"
    delete_on_termination = true
  }
}

#
# EIP ASSOCIATION
#
data "aws_eip" "runtime_eip" {
  filter {
    name   = "tag:Name"
    values = ["runtime-ephemeral"]
  }
}

resource "aws_eip_association" "runtime_eip_association" {
  instance_id   = aws_instance.runtime.id
  allocation_id = data.aws_eip.runtime_eip.id
}

#
# IAM ROLE
#
resource "aws_iam_role" "runtime_role" {
  name = "runtime-ephemeral-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  permissions_boundary = var.permissions_boundary_arn
}

resource "aws_iam_instance_profile" "runtime_profile" {
  name = "runtime-ephemeral-EC2InstanceProfile"
  role = aws_iam_role.runtime_role.name
}

#
# SSM POLICY
#
data "aws_iam_policy" "AmazonSSMManagedInstanceCore" {
  arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_iam_role_policy_attachment" "ssm_policy" {
  role       = aws_iam_role.runtime_role.name
  policy_arn = data.aws_iam_policy.AmazonSSMManagedInstanceCore.arn
}

#
# SECRETS MANAGER POLICY
#
resource "aws_iam_policy" "secrets_manager_policy" {
  name        = "runtime-ephemeral-SecretsManagerPolicy"
  description = "Allows access to specific secrets"

  # Define the policy JSON
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:ListSecrets",
          "secretsmanager:DescribeSecret"
        ]
        Resource = "arn:aws:secretsmanager:${var.region}:${data.aws_caller_identity.current.account_id}:secret:runtime-tls-*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "secrets_manager_policy_attachment" {
  role       = aws_iam_role.runtime_role.name
  policy_arn = aws_iam_policy.secrets_manager_policy.arn
}

#
# SECURITY GROUP
#
resource "aws_security_group" "security_group" {
  name = "runtime-ephemeral-sg-${random_id.unique_id.hex}"
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Allow HTTP traffic but istio will redirect to HTTPS
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # for local testing
  dynamic "ingress" {
    for_each = var.enable_ssh ? [1] : []
    content {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      cidr_blocks = ["${var.ssh_ip}/32"]
    }
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

resource "aws_key_pair" "ssh" {
  count = var.enable_ssh ? 1 : 0

  key_name   = "runtime-dev-key"
  public_key = tls_private_key.ssh[0].public_key_openssh
}
