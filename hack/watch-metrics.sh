#!/bin/sh

# Copyright 2024 Defense Unicorns
# SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

NAMESPACE="uds-runtime"
INTERVAL=1 # in seconds

while true; do
  kubectl top pods --namespace=$NAMESPACE
  sleep $INTERVAL
done
