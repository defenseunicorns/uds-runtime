# Runtime Ephemeral Infrastructure

The UDS Runtime IAC is used by the [nightly-infra workflow](../workflows/nightly-infra.yaml), via [uds tasks](./tasks/infra.yaml), to destroy and create ephemeral testing clusters, using the latest `nightly-unstable` image of UDS Runtime.

## How it Works

When the nightly workflow kicks off, it will `init` using the backend variables defined in the workflow to then destroy the currently running ec2 instance and cluster. After removing the old instance, it will create a new ec2 instance in the UDS CI AWS account, that on startup will do the following:

1. clone the [uds-k3d](https://github.com/defenseunicorns/uds-k3d) repo, setting `nginx.conf` to redirect for the `.burning.boats` domain
1. run the default task of `uds-k3d`, creating the k3d cluster on the instance
1. setup the `kubecontext` to be used by `uds`
1. pull the `.burning.boats` tls cert and key from secrets manager
1. deploy the `init` and `UDS Core` packages
1. deploy the `UDS Runtime` package

## Custom AMI

The ec2 instance is created with a custom AMI. We use `packer` to define the AMI in [runtime.pkr.hcl](./packer/runtime.pkr.hcl) and build / push it to our AWS accounts.

## Development and Testing

> **NOTE**  
> **Please use the UDS Dev AWS Account instead of CI**

For local development and testing:

pre-requisites:
* [opentofu](https://opentofu.org/)
* [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

1. Make sure you're terminal is authenticated to the AWS Dev account
1. Create a state bucket and dynamo table (either via CLI or through UI)
1. Alter the [variables](./terraform/variables.tf)
    * set the region to `us-east-1`
    * set the permissions boundary arn / name. You can find that under policies in the IAM console.
    * If you want to debug using SSH -- enable ssh and add your public IP.
1. Comment out the EIP association in [main.tf](./terraform//main.tf). This EIP is a dedicated EIP in the CI account attached to the `runtime-canary.burning.boats` domain.
1. Init and Apply:

    Via uds task from the root level of this repo:  `uds run -f .github/test-infra/tasks/infra.yaml create-iac`

    OR:

    ```bash
    cd .github/test-infra/terraform
    tofu init
    tofu apply -auto-approve
    ```


> **WARNING**  
> **DO NOT PUSH CHANGES TO VARIABLES SUCH AS ENABLING SSH AND PERMISSIONS BOUNDARY INFORMATION**

## Debug with SSH

If you enabled ssh and added your IP when developing locally, you can access your instance using the `runtime-dev.pem` that gets dropped in `.github/test-infra/terraform`.

```bash
ssh -i /path/to/runtime-dev.pem ubuntu@<public-ip>
```

## Debug with SSM

The ec2 instance has been configured with SSM for debugging running clusters without needing SSH. To start an SSM session:

 `Systems Manager` > click `Session Manager` under `Node Management` > click `start session` > select `runtime-ephemeral-*` > click `start session`
