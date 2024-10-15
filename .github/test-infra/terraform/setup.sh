#!/bin/sh

# Copyright 2024 Defense Unicorns
# SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

# clone and alter uds-k3d with new domain
git clone https://github.com/defenseunicorns/uds-k3d.git

# Define the file path
file_path="uds-k3d/chart/templates/nginx.yaml"

# # Replace 'uds.dev' with 'exploding.boats'
sed -i 's/uds\.dev/burning.boats/g' "$file_path"

# # Deploy cluster
cd uds-k3d && uds run

# Add ssm-user to docker group
usermod -aG docker ssm-user

# Restart SSM agent to apply changes
systemctl restart amazon-ssm-agent

# Get kubeconfig
mkdir -p /home/ubuntu/.kube
k3d kubeconfig get uds > /home/ubuntu/.kube/config

# CD to home directory or uds can't find the kubeconfig
cd /home/ubuntu

# Get TLS cert and key
TLS_CERT=$(aws secretsmanager get-secret-value --secret-id "runtime-tls-cert" --query "SecretString" --output text --region us-west-2 | uds zarf tools yq --input-format json '."runtime-tls-cert"')
TLS_KEY=$(aws secretsmanager get-secret-value --secret-id "runtime-tls-key" --query "SecretString" --output text --region us-west-2 | uds zarf tools yq --input-format json '."runtime-tls-key"')

export UDS_ADMIN_TLS_CERT=$TLS_CERT
export UDS_ADMIN_TLS_KEY=$TLS_KEY

# Enable TLS 1.2 support for admin gateway to support route53 health checks:
# https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html
export UDS_ADMIN_TLS1_2_SUPPORT=true

export UDS_TENANT_TLS_CERT=$TLS_CERT
export UDS_TENANT_TLS_KEY=$TLS_KEY

# k3d-core-demo:latest >= 0.27.3 deploys runtime on admin gateway. deploying nightly unstable afterward overwrites and redeploys runtime on tenant gateway without authsvc
uds deploy k3d-core-demo:latest --packages=init,core --set DOMAIN=burning.boats --confirm
uds zarf package deploy oci://ghcr.io/defenseunicorns/packages/uds/uds-runtime:nightly-unstable --confirm
