# Local plugin execution

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This directory provides an example of generating code stubs from a [Buf input][input] using Protobuf plugins. Here, we'll use two locally installed plugins:

Plugin | Language
:------|:--------
[`protoc-gen-go`][protoc-gen-go] | [Go]
[`protoc-gen-rust`][protoc-gen-rust] | [Rust]

Once those executables are installed, you can use the [`buf` CLI][cli], in conjunction with the configuration in [`buf.gen.yaml`](./buf.gen.yaml), to generate code stubs from any valid [Buf input][input]. Let's try it with the [`buf.build/buf-examples/observabilityapi`][api] module from the [workspace] example:

```sh
buf generate buf.build/buf-examples/observabilityapi
```

One drawback of this local plugin approach is that you need to install the executables on your own, make sure that they're on your `PATH`, and so on. To see how to generate code stubs without needing to install _any_ plugins locally, see the [remote plugin execution][remote] example project.

[api]: https://buf.build/buf-examples/observabilityapi
[cli]: https://github.com/bufbuild/buf
[go]: https://golang.org
[input]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
[module]: https://buf.build/buf-examples/observabilityapi
[protoc-gen-go]: https://github.com/golang/protobuf
[protoc-gen-rust]: https://crates.io/crates/protobuf-codegen
[remote]: ../plugin-execution-remote
[rust]: https://rust-lang.org
[workspace]: ../workspace
