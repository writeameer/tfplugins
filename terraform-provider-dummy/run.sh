#!/usr/bin/env bash

export TF_LOG=info
unset TF_LOG
terraform init
terraform plan
terraform apply -auto-approve

