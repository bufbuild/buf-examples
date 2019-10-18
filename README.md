# Buf Example

[![Travis CI](https://img.shields.io/travis/com/bufbuild/buf-example/master)](https://travis-ci.com/bufbuild/buf-example)
[![GitHub Actions](https://github.com/bufbuild/buf-example/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-example/actions?workflow=CI)

This is a simple example of [Buf](github.com/bufbuild/buf) usage that:

- Installs `buf` from GitHub Releases.
- Runs linting and breaking change detection.

The default `make` target is `make local`, which compares against the head of your
local git master branch for breaking change detection.

The `make` target `make remote` compares against the head of the remote git
master branch, due to Travis not cloning any branch except for that under test.

For GitHub Actions, you can use `make local` with a trick that clones the branch
you want first, and then creates a local head for it. See [ci.yaml](.github/workflows/ci.yaml)
for an example.

See [buf.build](https://buf.build) for documentation.
