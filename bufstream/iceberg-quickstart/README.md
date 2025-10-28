![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Bufstream and Iceberg quickstart

This directory contains example code for Bufstream's [Iceberg quickstart][docs].
It walks through the key tasks for creating a local Bufstream and Iceberg environment:

1. Deploying a Docker-based Bufstream environment with local object storage and Iceberg REST catalog.
2. Configuring Bufstream and a topic for Iceberg archival.
3. Running the Iceberg archival process on demand.
4. Querying Iceberg with Apache Spark and a Jupyter Notebook.

Note that the MinIO image within this Docker Compose project is provided by [lithus](https://github.com/golithus/minio-builds). [MinIO ceased publishing public Docker images](https://github.com/minio/minio/issues/21647) in October 2025.


[docs]: https://buf.build/docs/bufstream/iceberg/quickstart/
