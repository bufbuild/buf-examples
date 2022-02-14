# Buf on GitHub Actions example

[![GitHub Actions](https://github.com/bufbuild/buf-example/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-example/actions?workflow=CI)

> The configuration for [GitHub Actions][actions] is in the [`.github`](../.github) directory in the root, as that's where GitHub expects it to be.

This example uses three Buf-specific GitHub Actions:

* [`buf-setup-action`][buf-setup] installs the [`buf` CLI][cli]
* [`buf-lint-action`][buf-lint] [lints][lint] this Protobuf module
* [`buf-breaking-action`][buf-breaking] runs [breaking change detection][breaking] against the current `main` branch

[actions]: https://docs.github.com/actions
[breaking]: https://docs.buf.build/breaking
[buf-breaking]: https://github.com/bufbuildc/buf-breaking-action
[buf-lint]: https://github.com/bufbuildc/buf-lint-action
[buf-setup]: https://github.com/bufbuildc/buf-setup-action
[cli]: https://github.com/bufbuild/buf
[lint]: https://docs.buf.build/lint
