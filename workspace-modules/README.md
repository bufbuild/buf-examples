# Buf workspace example (modules only)

This directory provides an example of a Buf [workspace] that only contains Buf [modules] intended to be consumed as dependencies by other Protobuf projects. That means that there's no [`buf.gen.yaml`] [buf-gen-yaml] configuration for [generating code stubs][generate].

## Structure

```
├── buf.work.yaml
├── observabilityapi
│   ├── api
│   │   ├── v1
│   │   │   └── api.proto
│   │   └── v2
│   │       └── api.proto
│   └── buf.yaml
└── observabilitytypes
    ├── buf.yaml
    ├── log
    │   ├── v1
    │   │   └── log.proto
    │   ├── v2
    │   │   └── log.proto
    │   └── v3
    │       └── log.proto
    └── metric
        └── v1
            └── metric.proto
```

The workspace configuration is in [`buf.work.yaml`](./buf.work.yaml).

## Tasks

To [lint] all of the modules inside the workspace:

```sh
buf lint
```

The `buf` CLI lints _all_ of the modules in a workspace by default. To lint a specific module:

```sh
buf lint observabilityapi
```

To perform [breaking change detection][breaking] against the `main` branch of this repository:

```sh
buf breaking \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace-modules
```

As with `buf lint`, `buf breaking` detects breaking changes amongst all modules. To target only a specific module:

```sh
buf breaking observabilitytypes \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace-modules
```

[breaking]: https://docs.buf.build/breaking
[generate]: https://docs.buf.build/generate
[lint]: https://docs.buf.build/lint
[modules]: https://docs.buf.build/bsr/overview#module
[workspace]: https://docs.buf.build/reference/workspaces
