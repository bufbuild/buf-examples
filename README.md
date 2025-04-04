![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Buf examples

This repository houses example projects and code for products in the [Buf][buf] ecosystem. For more information, we recommend checking out the [official docs][docs].

## Buf CLI

[This directory](./cli) contains quickstart and example code that demonstrates key functions and usage of the [Buf CLI](https://github.com/bufbuild/buf).

## Buf Schema Registry (BSR)

[This directory](./bsr) contains quickstart code that demonstrates key functions and usage of the [BSR][bsr], and how to create [Buf plugins][plugins] for custom lint and breaking change rules that you can upload to the BSR.

## Protovalidate

[Quickstart guides and examples](./protovalidate) for [Protovalidate][protovalidate] show how standard annotations and custom CEL expressions can be used to validate Protobuf messages.

## Integrations

[Integrations](./integrations) contains example code supporting Buf's documentation about integration with CI/CD pipelines.

[bsr-docs]: https://buf.build/docs/bsr/
[bsr]: https://buf.build
[buf]: https://github.com/bufbuild
[cli-docs]: https://buf.build/docs/cli/
[docs]: https://buf.build/docs
[plugins]: https://buf.build/docs/cli/buf-plugins/overview/
[protovalidate]: https://buf.build/docs/protovalidate/