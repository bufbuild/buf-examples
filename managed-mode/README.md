# Managed mode

This example shows Buf's [managed mode][mm] in practice. Managed mode is a configuration option that
you can set in your [`buf.gen.yaml`][buf-gen-yaml] configuration file. If you set it to `true`, it
instructs the [`buf` CLI][repo] to use an opinionated set of values for each of `protoc`'s
natively supported languages, such as Go, Java, and C#. These options are written on the fly into
your Protobuf sources so that you don't need to include them in your `.proto` files directly.

[buf-gen-yaml]: https://docs.buf.build/configuration/v1/buf-gen-yaml#managed
[mm]: https://docs.buf.build/generate/managed-mode
[repo]: https://github.com/bufbuild/buf
