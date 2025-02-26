FROM golang:1.22-bookworm AS builder
WORKDIR /opt
COPY ../.. /opt
ARG access_token
RUN \
  apt-get update \
  && apt-get install ca-certificates
RUN make build

FROM debian:bookworm-slim
WORKDIR /opt
COPY --from=builder /opt/ /opt
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["/opt/gncompany"]
