# Buf workspace example

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This directory provides an example of a Buf [workspace] that contains multiple Buf [modules] that are meant to be consumed as dependencies by other Protobuf projects. There are two modules in this workspace:

* [`observabilityapi`](./observabilityapi) defines an observability-related service.
* [`observabilitytypes`](./observabilitytypes) defines many of the types used by the service in the `observabilityapi` module.

Both modules are available on the [Buf Schema Registry][bsr]:

* [`buf.build/buf-examples/observabilityapi`][bsr-api]
* [`buf.build/buf-examples/observabilitytypes`][bsr-types]

## Structure

The basic structure of this workspace:

```
.
├── buf.gen.yaml
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

## Linting

To [lint] all of the Buf modules in the workspace:

```sh
buf lint
```

To lint specific modules:

```sh
buf lint observabilityapi
buf lint observabilitytypes
```

## Breaking change detection

To perform [breaking change detection][breaking] against the previous commit to the `main` branch of this repository:

```sh
buf breaking \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace
```

As with `buf lint`, `buf breaking` detects breaking changes amongst all modules. To target only a specific module:

```sh
buf breaking observabilitytypes \
--against ../.git#branch=main,ref=HEAD~1,subdir=workspace
```

In this case, the `--against` flag targets the root of this repository using `../.git` and then this directory using `subdir=workspace`. In most projects, you would likely be able to use a more straightforward target, like this:

```sh
buf breaking observabilitytypes \
--against .git#branch=main,ref=HEAD~1
```

## Code generation

To [generate code stubs][generate] for all modules in the workspace using the plugins defined in [`buf.gen.yaml`](./buf.gen.yaml):

```sh
buf generate
```

To generate code stubs for specific modules:

```sh
buf generate observabilityapi
buf generate observabilitytypes
```

[breaking]: https://docs.buf.build/breaking
[bsr]: https://docs.buf.build/bsr
[bsr-api]: https://buf.build/buf-examples/observabilityapi
[bsr-types]: https://buf.build/buf-examples/observabilitytypes
[buf-work-yaml]: https://docs.buf.build/configuration/v1/buf-work-yaml
[generate]: https://docs.buf.build/generate
[install]: https://docs.buf.build/installation
[lint]: https://docs.buf.build/lint
[modules]: https://docs.buf.build/bsr/overview#module
[workspace]: https://docs.buf.build/reference/workspaces
