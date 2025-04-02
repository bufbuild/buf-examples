![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Buf CLI examples

The [Buf CLI][cli] is a crucial tool for managing Protobuf schemas. This directory supplements its [documentation][documentation] with examples showing how to use its features for code generation, breaking change detection, and linting tasks.  

## Workspaces

[Example code](./workspace) for the [Buf CLI][cli] provides an example of using one `buf.gen.yaml` file to manage multiple modules of Protobuf files in one repository.

## Local plugin execution

[Example code](./plugin-execution-local) for [local plugin execution][plugin] provides an example of code generation using local `protoc` plugins. 

## Remote plugin execution

[Example code](./plugin-execution-remote) for [remote plugin execution][remote] provides an example of code generation using plugins hosted on the Buf Schema Registry (BSR).

## Managed mode

[Example code](./managed-mode) for [generating code stubs using managed mode][managed] provides an example of controlling file and field options without adding language-specific options to Protobuf files.

## Breaking change detection

[Example code](./breaking-change-detection) for [breaking change detection][breaking] provides an example of using Buf's breaking change detection to ensure that your organization can evolve Protobuf schemas quickly and safely.

## Linting

[Example code](./linting) for [linting][lint] provides both "good" and "bad" examples of Protobuf files that do and do not meet linting rules enforced by the Buf CLI. 

## Limiting output types

[Example code](./limit-output-types) for the [buf build][limit-types] command showing how to limit output types with `--type`.

[cli]: https://github.com/bufbuild/buf
[documentation]: https://buf.build/docs/cli/
[lint]: https://buf.build/docs/lint/overview/
[breaking]: https://buf.build/docs/breaking/overview/
[remote]: https://buf.build/docs/generate/tutorial/#generate-code-using-remote-plugins
[plugin]: https://buf.build/docs/generate/tutorial/#generate-code-using-local-plugins
[managed]: https://buf.build/docs/generate/managed-mode/
[limit-types]: https://buf.build/docs/reference/cli/buf/build/#flags


[actions]: https://docs.github.com/actions
[bsr]: https://docs.buf.build/bsr
[bsr-org]: https://buf.build/buf-examples
[buf]: https://buf.build
[ci]: https://docs.buf.build/ci-cd
[cli]: https://buf.build/docs/cli/quickstart/
[docs]: https://docs.buf.build
[modules]: https://docs.buf.build/bsr/overview#modules
[min-version]: https://github.com/bufbuild/buf/releases/tag/v1.32.0
[v2-buf-yaml]: https://buf.build/docs/configuration/v2/buf-yaml
[v2-buf-gen-yaml]: https://buf.build/docs/configuration/v2/buf-gen-yaml

