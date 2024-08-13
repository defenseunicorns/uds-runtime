#!/bin/sh

# clone and alter uds-k3d with new domain
git clone https://github.com/defenseunicorns/uds-k3d.git

# Define the file path
 file_path="uds-k3d/chart/templates/nginx.yaml"

# # Replace 'uds.dev' with 'exploding.boats'
 sed -i 's/uds\.dev/exploding.boats/g' "$file_path"

# # Deploy cluster
cd uds-k3d && uds run default

cd ..

# Get kubeconfig for uds to find
mkdir -p /home/ubuntu/.kube
k3d kubeconfig get uds > /home/ubuntu/.kube/config

# Get TLS cert and key
CA_CERT=$(aws ssm get-parameter --name "runtime-ephemeral-ca-cert" --with-decryption --query "Parameter.Value" --output text)
TLS_KEY=$(aws ssm get-parameter --name "runtime-ephemeral-key" --with-decryption --query "Parameter.Value" --output text)
TLS_CERT=$(aws ssm get-parameter --name "runtime-ephemeral-cert" --with-decryption --query "Parameter.Value" --output text)

export UDS_CA_CERT=$CA_CERT
export UDS_ADMIN_TLS_CERT=$TLS_CERT
export UDS_ADMIN_TLS_KEY=$TLS_KEY
export UDS_TENANT_TLS_CERT=$TLS_CERT
export UDS_TENANT_TLS_KEY=$TLS_KEY

# CD to home directory or uds can't find the kubeconfig
cd /home/ubuntu
uds zarf tools kubectl config get-contexts

uds deploy ghcr.io/defenseunicorns/packages/uds/bundles/k3d-core-slim-dev:0.25.2 --packages=init,core-slim-dev --set DOMAIN=exploding.boats --confirm
uds zarf package deploy oci://ghcr.io/defenseunicorns/packages/uds/uds-runtime:nightly-unstable --confirm
