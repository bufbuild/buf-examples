# Formatting Protobuf sources

> You need to have the [`buf` CLI][install] installed to follow along with this example.

Buf enables you to [format] Protobuf sources using the `buf format` command. The formatter is _not_ configurable and uses only an opinionated set of rules.

```sh
buf format --diff ./bad
```

To fix:

```sh
buf format --write ./bad
```

[format]: https://docs.buf.build/format
[install]: https://docs.buf.build/installation
