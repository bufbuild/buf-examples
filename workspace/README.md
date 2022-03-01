# Buf workspace example

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This directory provides an example of a Buf [workspace] that contains multiple Buf [modules] that are intended to be consumed as dependencies by other Protobuf projects.

## Structure

```
.
├── buf.work.yaml
├── observabilityapi
│   ├── api
│   │   ├── v1
│   │   │   └── api.proto
│   │   └── v2
│   │       └── api.proto
│   └── buf.yaml
└── observabilitytypes
    ├── buf.yaml
    ├── log
    │   ├── v1
    │   │   └── log.proto
    │   ├── v2
    │   │   └── log.proto
    │   └── v3
    │       └── log.proto
    └── metric
        └── v1
            └── metric.proto
```

The [workspace configuration][buf-work-yaml] is defined in [`buf.work.yaml`](./buf.work.yaml).

## Tasks

To [lint] all of the Buf modules in the workspace:

```sh
buf lint
```

To lint specific modules:

```sh
buf lint observabilityapi
buf lint observabilitytypes
```

To perform [breaking change detection][breaking] against the `main` branch of this repository:

```sh
buf breaking \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace
```

As with `buf lint`, `buf breaking` detects breaking changes amongst all modules. To target only a specific module:

```sh
buf breaking observabilitytypes \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace-modules
```

[breaking]: https://docs.buf.build/breaking
[buf-work-yaml]: https://docs.buf.build/configuration/v1/buf-work-yaml
[generate]: https://docs.buf.build/generate
[install]: https://docs.buf.build/installation
[lint]: https://docs.buf.build/lint
[modules]: https://docs.buf.build/bsr/overview#module
[workspace]: https://docs.buf.build/reference/workspaces
