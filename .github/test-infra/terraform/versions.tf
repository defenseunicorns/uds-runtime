# Copyright 2024 Defense Unicorns
# SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

terraform {
  backend "s3" {}
  required_version = ">= 1.6.0, <= 1.8.1"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.62.0"
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 3.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = ">= 0.9.1"
    }
  }
}
