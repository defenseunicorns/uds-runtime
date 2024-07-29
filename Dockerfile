FROM cgr.dev/chainguard/static:latest

# grab auto platform arg
ARG TARGETARCH

# 65532 is the UID of the `nonroot` user in chainguard/static.
# See: https://edu.chainguard.dev/chainguard/chainguard-images/reference/static/overview/#users
USER 65532:65532

# copy binary from local and expose port
COPY --chown=65532:65532 build/uds-runtime-${TARGETARCH} /app/uds-runtime
ENV PORT=8080
EXPOSE 8080

# run binary
CMD ["./app/uds-runtime"]
