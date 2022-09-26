# Buf on GitHub Actions example

[![GitHub Actions](https://github.com/bufbuild/buf-examples/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-examples/actions?workflow=CI)

This example project shows you how to use Buf in a [GitHub Actions][actions] setting. The Actions pipeline here involves three Buf-specific GitHub Actions:

* [`buf-setup-action`][buf-setup] installs the [`buf` CLI][cli]
* [`buf-lint-action`][buf-lint] [lints][lint] this Protobuf module
* [`buf-breaking-action`][buf-breaking] runs [breaking change detection][breaking] against the current `main` branch

> The configuration for [GitHub Actions][actions] is in the [`.github`](../.github) directory in the root, as that's where GitHub expects it to be.

[actions]: https://docs.github.com/actions
[breaking]: https://docs.buf.build/breaking
[buf-breaking]: https://github.com/bufbuild/buf-breaking-action
[buf-lint]: https://github.com/bufbuild/buf-lint-action
[buf-setup]: https://github.com/bufbuild/buf-setup-action
[cli]: https://github.com/bufbuild/buf
[lint]: https://docs.buf.build/lint
