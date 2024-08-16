#!/bin/sh

# clone and alter uds-k3d with new domain
git clone https://github.com/defenseunicorns/uds-k3d.git

# Define the file path
 file_path="uds-k3d/chart/templates/nginx.yaml"

# # Replace 'uds.dev' with 'exploding.boats'
 sed -i 's/uds\.dev/burning.boats/g' "$file_path"

# # Deploy cluster
cd uds-k3d && uds run

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
export UDS_TENANT_TLS_CERT=$TLS_CERT
export UDS_TENANT_TLS_KEY=$TLS_KEY

uds deploy ghcr.io/defenseunicorns/packages/uds/bundles/k3d-core-demo:0.25.2 --packages=init,core --set DOMAIN=burning.boats --confirm
uds zarf package deploy oci://ghcr.io/defenseunicorns/packages/uds/uds-runtime:nightly-unstable --confirm
