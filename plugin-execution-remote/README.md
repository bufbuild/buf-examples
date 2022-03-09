# Remote plugin execution

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This directory provides an example of generating code stubs from a [Buf input][input] using Protobuf plugins. Here, we'll use two remote plugins hosted on the [Buf Schema Registry][bsr]:

Plugin | Language
:------|:--------
[`protocolbuffers/go`][pb-go] | [Go]
[`protocolbuffers/java`][pb-java] | Java

You can use the [`buf` CLI][cli], in conjunction with the configuration in [`buf.gen.yaml`](./buf.gen.yaml), to generate code stubs from any valid [Buf input][input]. Let's try it with the [`buf.build/buf-examples/observabilityapi`][api] module from the [workspace] example:

```sh
buf generate buf.build/buf-examples/observabilityapi
```

The [remote plugin execution][blog] was designed to avoid the pitfalls of using [locally installed plugins][local].

[api]: https://buf.build/buf-examples/observabilityapi
[blog]: https://buf.build/blog/remote-plugin-execution
[bsr]: https://docs.buf.build/bsr
[cli]: https://github.com/bufbuild/buf
[go]: https://golang.org
[input]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
[local]: ../plugin-execution-local
[module]: https://buf.build/buf-examples/observabilityapi
[pb-go]: https://buf.build/protocolbuffers/plugins/go
[pb-java]: https://buf.build/protocolbuffers/plugins/java
[rust]: https://rust-lang.org
[workspace]: ../workspace
