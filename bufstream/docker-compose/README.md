![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Docker Compose for Bufstream

This directory contains an example `docker-compose.yml` file that sets up all required infrastructure for a [Bufstream][bufstream] environment suitable for local testing and development.

## Prerequisites

* [Docker][docker] must be installed. 
* If you're using Docker Desktop, [host mode networking][docker-host-mode-networking] must be enabled.

## Starting the project

Just run:

```
$ docker compose up
```

## Services

This Compose project will start the following services:

- MinIO (https://min.io/) for S3 compatible storage.
  * Its API endpoint is available at http://localhost:9000.
  * Its admin endpoint is available at http://localhost:9001.
- A MinIO Client (mc - https://min.io/docs/minio/linux/reference/minio-mc.html) to bootstrap an initial bucket.
- PostgreSQL (https://www.postgresql.org/) for metadata storage.
  * It is available on the standard Postgres port (5432).
- A Bufstream broker.
  * Its Kafka endpoint is available at localhost:9092.
  * Its admin endpoint listens on its default port (9089), allowing you to run [`bufstream admin`](https://buf.build/docs/bufstream/reference/cli/admin/) commands against the broker.

## Memory

The Compose project uses the `BUFSTREAM_AVAILABLE_MEMORY` environment variable to limit Bufstream to 1 GiB of memory.
This should be suitable for simple local development and testing, but may be insufficient for high-throughput workloads or performance testing.

If you encounter out-of-memory issues, increase the value in the `docker-compose.yml` file. For reference, production Kubernetes deployments default to 8 GiB (8589934592 bytes).

## Volumes

When this Compose project starts, it will create two directories:

1. `minio`: A volume used by the `minio` service for object storage.
2. `postgres`: A volume used by the `postgres` service for its database storage.

Before committing this `docker-compose.yml` file to your own repository, add these to your `.gitignore`.

## Networks

This example uses host networking. You may encounter issues if there are port conflicts. If you need to use port mappings, see [Docker's documentation](https://docs.docker.com/compose/how-tos/networking/).

[bufstream]: https://buf.build/product/bufstream
[docker]: https://docs.docker.com/engine/install/
[docker-host-mode-networking]: https://docs.docker.com/engine/network/drivers/host/#docker-desktop
