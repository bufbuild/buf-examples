![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Buf Schema Registry examples

The Buf Schema Registry (BSR) is the missing package manager for Protobuf, allowing you to manage schemas, dependencies, and governance at scale.

Supplementing its [documentation], this repository provides quickstart code exercises and reference examples for its use.

## Buf Schema Registry quickstart

[Quickstart code](./quickstart) for the Buf Schema Registry's [quickstart][quickstart] introduces you to key features of the BSR, including managing dependencies, creating and publishing documented modules, and using generated SDKs. It contains a `start` directory with a starting state and `finish` with a working solution.

## Buf plugins quickstart

[Quickstart code](./buf-check-plugin) for the Buf plugins [quickstart][plugins] walks through how to create a Buf plugin for custom lint and breaking change rules. It contains a `start` directory with a starting state and `finish` with a working solution.

[documentation]: https://buf.build/docs/bsr
[plugins]: https://buf.build/docs/cli/buf-plugins/tutorial-create-buf-plugin/
[quickstart]: https://buf.build/docs/bsr/quickstart/
