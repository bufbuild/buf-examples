![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Docker Compose for Bufstream Quickstart

This directory contains an example docker compose that sets up all the requirement
infrastrucutre to have a playground environment for Bufstream.

It sets up:
- MinIO (https://min.io/) for S3 compatible storage
  * Its API endpoint is made available at http://localhost:9000
  * Its admin endpoint is made available at http://localhost:9001
- A MinIO Client (mc - https://min.io/docs/minio/linux/reference/minio-mc.html) for creating a MinIO bucket.
- PostgreSQL (https://www.postgresql.org/) for metadata storage
  * Its API endpoint is made available at http://localhost:5432
- Bufstream itself
  * Its Kafka endpoint is made available at localhost:9092
  * Its Admin endpoint is made available at a Connect endpoint on http://localhost:9089

Simply run:

```
$ docker compose up
```
