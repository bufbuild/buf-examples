# Buf Example

[![Travis CI](https://img.shields.io/travis/com/bufbuild/buf-example/master)](https://travis-ci.com/bufbuild/buf-example)
[![CircleCI](https://img.shields.io/circleci/build/github/bufbuild/buf-example/master)](https://circleci.com/gh/bufbuild/buf-example)
[![GitHub Actions](https://github.com/bufbuild/buf-example/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-example/actions?workflow=CI)

This is a simple example of [Buf](github.com/bufbuild/buf) usage that:

- Installs `buf` from GitHub Releases.
- Runs linting and breaking change detection.

CI is set up for:

- [Travis CI](https://travis-ci.com/bufbuild/buf-example)
- [CircleCI](https://circleci.com/gh/bufbuild/buf-example)
- [GitHub Actions](https://github.com/bufbuild/buf-example/actions?workflow=CI)

The default `make` target is `make local`, which compares against the head of your
local git master branch for breaking change detection.

The `make` targets `make https` and `make ssh` compare against the head of the remote git
master branch, due to Travis and CircleCI not cloning any branch except for that under test.
In this example, we use `https` for Travis, and both `https` and `ssh` for CircleCI, but you only
need to choose one of `https` and `ssh`.

For GitHub Actions, you can use `make local` with a trick that clones the branch
you want first, and then creates a local head for it. See [ci.yaml](.github/workflows/ci.yaml)
for an example.

See [buf.build](https://buf.build) for documentation.
