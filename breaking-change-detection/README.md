# Breaking change detection

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This project shows Buf [breaking change detection][breaking] in action. There are three different [Buf modules][modules] in play here:

* An [`initial`](./initial) module
* A module called [`compatible`](./compatible) that introduces non-breaking changes to `initial`
* A module called [`incompatible`](./incompatible) that introduces several breaking changes to `initial`

To verify that `compatible` introduces no breaking changes against `initial`:

```sh
buf breaking ./initial --against ./compatible
```

No output indicates that breaking change detection was successful.

To verify that `incompatible` introduces breaking changes against `initial`:

```sh
buf breaking ./initial --against ./incompatible
```

That command provides granular insight into those breaking changes:

```
TODO
```

[breaking]: https://docs.buf.build
[install]: https://docs.buf.build/installation
[modules]: https://docs.buf.build/bsr/overview#modules
