# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

FROM cassandra:5.0@sha256:8819d1b7877ee6621751a761cc163332e9c3eaf671c8ef672f7c77a7c5bff05f

COPY schema/* /cassandra-schema/

ENV CQLSH_HOST=cassandra

RUN groupadd -g 65532 nonroot && \
    useradd -u 65532 -g nonroot nonroot --create-home

USER 65532:65532
ENTRYPOINT ["/cassandra-schema/docker.sh"]
