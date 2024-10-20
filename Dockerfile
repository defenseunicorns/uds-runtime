# Copyright 2024 Defense Unicorns
# SPDX-License-Identifier: AGPL-3.0-or-later OR LicenseRef-Defense-Unicorns-Commercial

FROM cgr.dev/chainguard/static:latest

# grab auto platform arg
ARG TARGETARCH

# 65532 is the UID of the `nonroot` user in chainguard/static.
# See: https://edu.chainguard.dev/chainguard/chainguard-images/reference/static/overview/#users
USER 65532:65532

# copy binary from local and expose port
COPY --chown=65532:65532 build/uds-runtime-linux-${TARGETARCH} /app/uds-runtime
ENV PORT=8080
ENV LOCAL_AUTH_ENABLED=false
EXPOSE 8080

# run binary
CMD ["./app/uds-runtime"]
