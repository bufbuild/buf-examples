![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Buf CLI examples

The [Buf CLI][cli] is a crucial tool for managing Protobuf schemas. This directory supplements its [documentation][documentation] with examples and quickstart code showing how to use its features for code generation, breaking change detection, and linting tasks.  

## Buf CLI quickstart

[Quickstart code](./quickstart) for the Buf CLI's [Quickstart][cli-quickstart] introduces you to key features of the Buf CLI, including setting up a local workspace, generating code, breaking change detection, and linting. It contains a `start` directory with a starting state and `finish` directory with a working solution.

## Breaking change detection

[Quickstart code](./breaking-change-detection) for the [breaking change detection quickstart][breaking] provides a walkthrough of using Buf's breaking change detection to ensure that your organization can evolve Protobuf schemas quickly and safely. It contains a `start` directory with a starting state and `finish` directory with a working solution.

## Linting

[Quickstart code](./linting) for the [lint quickstart][lint] provides a walkthrough of checking for lint errors, customizing your lint configuration, and temporarily ignoring errors. It contains a `start` directory with a starting state and `finish` directory with a working solution.

## Code generation

### Local plugin execution

[Example code](./plugin-execution-local) for [local plugin execution][plugin] provides an example of code generation using local `protoc` plugins. 

### Remote plugin execution

[Example code](./plugin-execution-remote) for [remote plugin execution][remote] provides an example of code generation using plugins hosted on the Buf Schema Registry (BSR).

### Managed mode

[Example code](./managed-mode) for [generating code stubs using managed mode][managed] provides an example of controlling file and field options without adding language-specific options to Protobuf files.

### Limiting output types

[Example code](./limit-output-types) for the [buf build][limit-types] command showing how to limit output types with `--type`.

## Workspaces

[Example code](./workspace) for the [Buf CLI][cli] provides an example of using one `buf.gen.yaml` file to manage multiple modules of Protobuf files in one repository.

[breaking]: https://buf.build/docs/breaking/quickstart/
[cli]: https://github.com/bufbuild/buf
[cli-docs]: https://buf.build/docs/cli/
[cli-quickstart]: https://buf.build/docs/cli/quickstart/
[limit-types]: https://buf.build/docs/reference/cli/buf/build/#flags
[lint]: https://buf.build/docs/lint/quickstart/
[managed]: https://buf.build/docs/generate/managed-mode/
[plugin]: https://buf.build/docs/generate/tutorial/#generate-code-using-local-plugins
[remote]: https://buf.build/docs/generate/tutorial/#generate-code-using-remote-plugins