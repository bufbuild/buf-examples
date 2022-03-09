# Local plugin execution

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This directory provides an example of generating code stubs. Here, we'll use two locally installed
plugins:

Plugin | Language
:------|:--------
[`protoc-gen-go`][protoc-gen-go] | [Go]
[`protoc-gen-rust`][protoc-gen-rust] | [Rust]

To install those plugins:

```sh
# TODO
```

Once those plugins are installed, you can generate code stubs from any valid [Buf input][input].

```sh
buf generate buf.build/buf-examples/observabilityapi
```

Borrowed from the [workspace] example.

[go]: https://golang.org
[input]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
[module]: https://buf.build/buf-examples/observabilityapi
[protoc-gen-go]: https://github.com/golang/protobuf
[protoc-gen-rust]: https://crates.io/crates/protobuf-codegen
[rust]: https://rust-lang.org
[workspace]: ../workspace
