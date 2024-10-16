# Buf example projects

This repo houses example projects for learning how to use [Buf]. For more information, we recommend checking out the [official docs][docs].

## Version

The examples in this repo use [v2 `buf.yaml`][v2-buf-yaml] and [v2 `buf.gen.yaml`][v2-buf-gen-yaml]
configurations. All examples require the `buf` CLI at version [v1.32.0][min-version] or higher.

## Projects

Project | Directory | Category
:-------|:----------|:--------
Buf [workspace] | [`workspace`](./workspace) | Code organization
[Linting][lint] Protobuf sources | [`linting`](./linting) | Maintenance
[Breaking change detection][breaking] | [`breaking-change-detection`](./breaking-change-detection) | Maintenance
[Remote plugin execution][remote] | [`plugin-execution-remote`](./plugin-execution-remote) | Code generation
[Local plugin execution][plugin] | [`plugin-execution-local`](./plugin-execution-local) | Code generation
Generating code stubs using [managed mode][managed] | [`managed-mode`](./managed-mode/) | Code generation
[Limit output types][limit-types] for `buf build` | [`limit-output-types`](./limit-output-types) | CLI
Buf on [GitHub Actions][actions] | [`github-actions`](./github-actions) | [CI/CD][ci]

Several of these projects available as modules on the [Buf Schema Registry][bsr] under the
[`buf-examples`][bsr-org] organization.

[actions]: https://docs.github.com/actions
[breaking]: https://docs.buf.build/breaking
[bsr]: https://docs.buf.build/bsr
[bsr-org]: https://buf.build/buf-examples
[buf]: https://buf.build
[ci]: https://docs.buf.build/ci-cd
[cli]: https://github.com/bufbuild/buf
[docs]: https://docs.buf.build
[limit-types]: https://docs.buf.build/build/usage#limit-to-specific-types
[lint]: https://docs.buf.build/lint
[managed]: https://docs.buf.build/generate/managed-mode
[modules]: https://docs.buf.build/bsr/overview#modules
[plugin]: https://docs.buf.build/bsr/remote-generation/concepts#plugins
[remote]: https://docs.buf.build/bsr/remote-generation/remote-plugin-execution
[workspace]: https://docs.buf.build/reference/workspaces
[min-version]: https://github.com/bufbuild/buf/releases/tag/v1.32.0
[v2-buf-yaml]: https://buf.build/docs/configuration/v2/buf-yaml
[v2-buf-gen-yaml]: https://buf.build/docs/configuration/v2/buf-gen-yaml
