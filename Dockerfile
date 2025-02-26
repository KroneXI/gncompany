FROM golang:1.22-bookworm AS builder
WORKDIR /opt
COPY ../.. /opt
ARG access_token
RUN \
  apt-get update \
  && apt-get install ca-certificates openssl
RUN make build

FROM debian:bookworm-slim
WORKDIR /opt
COPY --from=builder /opt/ /opt
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

RUN apt-get update && apt-get install -y openssl && \
  mkdir /opt/certs && \
  openssl req -x509 -newkey rsa:4096 -keyout /opt/certs/server.key -out /opt/certs/server.crt -days 365 -nodes -subj "/CN=localhost"

ENTRYPOINT ["/opt/gncompany"]
