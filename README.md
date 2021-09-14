# Buf Example

[![GitHub Actions](https://github.com/bufbuild/buf-example/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-example/actions?workflow=CI)
[![CircleCI](https://img.shields.io/circleci/build/github/bufbuild/buf-example/main)](https://circleci.com/gh/bufbuild/buf-example)

This is a simple example of [Buf](https://github.com/bufbuild/buf) usage that:

- Installs `buf` from GitHub Releases.
- Runs linting and breaking change detection.

CI is set up for:

- [GitHub Actions](https://github.com/bufbuild/buf-example/actions?workflow=CI)
- [CircleCI](https://circleci.com/gh/bufbuild/buf-example)

The default `make` target is `make local`, which compares against the head of your
local git main branch for breaking change detection.

The `make` targets `make https` and `make ssh` compare against the head of the remote git
main branch, due to CircleCI not cloning any branch except for that under test.
In this example, we use both `https` and `ssh` for CircleCI, but you only
need to choose one of `https` and `ssh`.

For GitHub Actions, you can use `make local` with a trick that clones the branch
you want first, and then creates a local head for it. See [ci.yaml](.github/workflows/ci.yaml)
for an example.

See [buf.build](https://buf.build) for documentation.
