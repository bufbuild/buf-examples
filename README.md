# Buf example projects

This repo houses example projects for learning how to use [Buf]. For more information, we recommend checking out the [official docs][docs].

## Version

The examples in this repo currently use version [1.1.0][version] of the [`buf` CLI][cli].

## Projects

Project | Directory | Category
:-------|:----------|:--------
Buf [workspace] | [`workspace-modules`](./workspace) | Code organization
[Linting][lint] Protobuf sources | [`linting`](./linting) | Maintenance
[Limit output types][limit-types] for `buf build` | [`limit-output-types`](./limit-output-types) | CLI
Buf on [CircleCI] | [`circleci`](./circleci) | [CI/CD][ci]
Buf on [GitHub Actions][actions] | [`github-actions`](./github-actions) | [CI/CD][ci]

[actions]: https://docs.github.com/actions
[buf]: https://buf.build
[ci]: https://docs.buf.build/ci-cd
[circleci]: https://circleci.com
[cli]: https://github.com/bufbuild/buf
[docs]: https://docs.buf.build
[limit-types]: https://docs.buf.build/build/usage#limit-to-specific-types
[lint]: https://docs.buf.build/lint
[modules]: https://docs.buf.build/bsr/overview#modules
[remote]: https://docs.buf.build/bsr/remote-generation/remote-plugin-execution
[version]: https://github.com/bufbuild/buf/releases/tag/v1.1.0
[workspace]: https://docs.buf.build/reference/workspaces
