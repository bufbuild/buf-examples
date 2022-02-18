# Managed mode

This example shows Buf's [managed mode][mm] in practice. Managed mode is a configuration option that you can set in your [`buf.gen.yaml`][buf-gen-yaml] configuration file. If you set `enabled` to `true`, it instructs the [`buf` CLI][repo] to use an opinionated set of values for each of `protoc`'s natively supported languages, such as Go, Java, and C#. These options are written on the fly into your Protobuf sources so that you don't need to include them in your `.proto` files directly.

To generate code stubs from the Protobuf sources in [`acme/weather/v1`](./acme/weather/v1):

```sh
buf generate
```

When you run this command, the `buf` CLI generates code stubs using the configuration in [`buf.gen.yaml`](./buf.gen.yaml). As you can see in that config, managed mode is enabled and several managed mode options is set. See the [`gen/proto`](./gen/proto) directory for generated code in these languages:

* [C++](./gen/proto/cpp)
* [C#](./gen/proto/csharp)
* [Go](./gen/proto/go)
* [Java](./gen/proto/java)
* [Objective-C](./gen/proto/objc)
* [PHP](./gen/proto/php)
* [Ruby](./gen/proto/ruby)

[arena]: https://developers.google.com/protocol-buffers/docs/reference/arenas
[buf-gen-yaml]: https://docs.buf.build/configuration/v1/buf-gen-yaml#managed
[mm]: https://docs.buf.build/generate/managed-mode
[repo]: https://github.com/bufbuild/buf
