#!/bin/sh

NAMESPACE="uds-runtime"
INTERVAL=1 # in seconds

while true; do
  kubectl top pods --namespace=$NAMESPACE
  sleep $INTERVAL
done
