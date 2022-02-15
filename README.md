# Buf example projects

This repo houses example projects for learning how to use [Buf]. For more information, we recommend checking out the [official docs][docs].

## Version

The examples in this repo currently use version [1.0.0-rc12][version] of the [`buf` CLI][cli].

## Projects

Project | Directory | Category
:-------|:----------|:--------
Buf on [CircleCI] | [`circleci`](./circleci) | [CI/CD][ci]
Buf on [GitHub Actions][actions] | [`github-actions`](./github-actions) | [CI/CD][ci]
Workspace with [Buf modules][modules] | [`workspace-modules`](./workspace-modules) | [Workspace]
Workspace with [remote plugin execution][remote] | [`workspace-remote-plugin-execution`](./workspace-remote-plugin-execution) | [Workspace]

[actions]: https://docs.github.com/actions
[buf]: https://buf.build
[ci]: https://docs.buf.build/ci-cd
[circleci]: https://circleci.com
[cli]: https://github.com/bufbuild/buf
[docs]: https://docs.buf.build
[modules]: https://docs.buf.build/bsr/overview#modules
[remote]: https://docs.buf.build/bsr/remote-generation/remote-plugin-execution
[version]: https://github.com/bufbuild/buf/releases/tag/v1.0.0-rc12
[workspace]: https://docs.buf.build/reference/workspaces
