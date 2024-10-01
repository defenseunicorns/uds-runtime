# Certs

The certs in this directory are primarily for used for dev'ing on UDS with HTTPS. They are also being used  to enable TLS when running UDS Runtime locally (such as when doing `uds ui`).

The certs are not sensitive and were taken from the UDS Core repo [here](https://github.com/defenseunicorns/uds-core/blob/main/src/istio/values/config-tenant.yaml); specifically these are the default certs for Istio tenant gateway.
