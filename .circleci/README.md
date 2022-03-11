# Buf on CircleCI example

[![CircleCI](https://img.shields.io/circleci/build/github/bufbuild/buf-examples/main)](https://circleci.com/gh/bufbuild/buf-examples)

This example project shows you how to use Buf in a [CircleCI][circle] setting. The workflow here involves two steps:

* [Linting][lint] the Protobuf module in the [`proto`](../proto) directory
* Running [breaking change detection][breaking] against the current `main` branch

The configuration for [CircleCI][circle] is in [`config.yml`](./config.yml).

## Other CI/CD examples

See the [`github`](../.github) project for another CI/CD example.

[breaking]: https://docs.buf.build/breaking
[circle]: https://circleci.com
[cli]: https://github.com/bufbuild/buf
[lint]: https://docs.buf.build/lint
