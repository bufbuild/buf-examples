![The Buf logo](https://raw.githubusercontent.com/bufbuild/buf-examples/main/.github/buf-logo.svg)

# Bufstream quickstart

This directory contains example code for Bufstream's basic quickstart.

It set up a minio for S3 storage and a postgres instance that bufstream connects to,
plus akhq for easy of exploration.

Simply run:

```
$ docker compose up
```

And once AKHQ is ready, navigate to http://localhost:8080/ to create and explore topics.

You can also access:

* minio's admin interface at http://localhost:9001/
* bufstream connect endpoint at http://localhost:8089/
* bufstream kafka endpoint at http://localhost:9092/

[docs]: https://buf.build/docs/bufstream/quickstart/
