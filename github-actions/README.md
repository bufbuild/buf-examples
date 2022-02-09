# Buf on GitHub Actions example

[![GitHub Actions](https://github.com/bufbuild/buf-example/workflows/CI/badge.svg)](https://github.com/bufbuild/buf-example/actions?workflow=CI)

This is a simple example of [Buf](https://github.com/bufbuild/buf) usage that:

- Installs `buf` from GitHub Releases.
- Runs linting and breaking change detection.

CI is set up for  [GitHub Actions](https://github.com/bufbuild/buf-example/actions?workflow=CI).

The default `make` target is `make local`, which compares against the head of your
local git main branch for breaking change detection.

For GitHub Actions, you can use `make local` with a trick that clones the branch
you want first, and then creates a local head for it. See [ci.yaml](.github/workflows/ci.yaml)
for an example.

See [buf.build](https://buf.build) for documentation.
