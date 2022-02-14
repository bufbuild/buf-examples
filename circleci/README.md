# Buf on CircleCI example

[![CircleCI](https://img.shields.io/circleci/build/github/bufbuild/buf-example/main)](https://circleci.com/gh/bufbuild/buf-example)

This example project shows you how to use Buf in a [CircleCI][circle] setting. The workflow here involves two steps:

* [Linting][lint] this Protobuf module
* Running [breaking change detection][breaking] against the current `main` branch

> The configuration for [CircleCI][circle] is in the [`.circleci`](../.circleci) directory in the root, as that's where CircleCI expects it to be.

## Other examples

See the [`github-actions`](../github-actions) project for another CI/CD example.

[breaking]: https://docs.buf.build/breaking
[circle]: https://circleci.com
[cli]: https://github.com/bufbuild/buf
[lint]: https://docs.buf.build/lint
