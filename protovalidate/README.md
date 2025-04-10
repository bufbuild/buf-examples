![The Buf logo](https://raw.githubusercontent.com/bufbuild/protovalidate/main/.github/buf-logo.svg)

# Protovalidate examples

[Protovalidate][protovalidate] provides standard annotations to validate common rules on messages and fields, as well as the ability to use [CEL][cel] to write custom rules. It's the next generation of [protoc-gen-validate][protoc-gen-validate], the only widely used validation library for Protobuf.

Supplementing its [documentation][protovalidate], this repository provides quickstart code exercises and reference examples for its use. 

## Protovalidate quickstart

Quickstart code for Protovalidate's [Quickstart][quickstart] is available in the following languages. Each contains a `start` directory with a starting state and `finish` with a working solution.

* [Go](quickstart-go/README.md)
* [Java](quickstart-java/README.md)
* [Python](quickstart-python/README.md)

## Compiling and generating code

Example code for [Adding Protovalidate to Protobuf projects][adding-protovalidate] is available for the following build tools:

* [The Buf CLI](compiling-buf/README.md)
* [protoc](compiling-protoc/README.md)

## Standard rules

[Example code](rules-standard/README.md) for [Using Protovalidate standard rules][standard-rules] provides a library of Protobuf files demonstrating Protovalidate's built-in standard rules.

## Custom rules

[Example code](rules-custom/README.md) for [Custom CEL rules][custom-rules] provides a library of Protobuf files demonstrating Common Expression Language (CEL)-based custom rules.

## Predefined rules

[Quickstart code](rules-predefined/README.md) for [Predefined rules][predefined-rules] shows how to create your own reusable rules. The only code changed during the tutorial is in Protobuf files, so no additional language examples are necessary. It contains a `start` directory with a starting state and `finish` with a working solution.

## Using Protovalidate in RPC APIs

Quickstart code showing how to use Protovalidate in unary RPCs is available in the following languages. Each contains a `start` directory with a starting state and `finish` with a working solution.

* [Quickstart code](connect-go/README.md) for [Protovalidate in Connect and Go][connect-go]
* [Quickstart code](grpc-go/README.md) for [Protovalidate in gRPC and Go][grpc-go]
* [Quickstart code](grpc-java/README.md) for [Protovalidate in gRPC and Java][grpc-java]
* [Quickstart code](grpc-python/README.md) for [Protovalidate in gRPC and Python][grpc-python]

## Using Protovalidate in Kafka

[Quickstart code](bufstream/README.md) for [Using Protovalidate in Kafka with Bufstream][bufstream] shows how to enforce schemas within the Bufstream Kafka broker. The only code changed during the tutorial is in Protobuf and YAML files, so only a Go example is provided. It contains a `start` directory with a starting state and `finish` with a working solution.

[protovalidate]: https://buf.build/docs/protovalidate/
[quickstart]: https://buf.build/docs/protovalidate/quickstart/
[adding-protovalidate]: https://buf.build/docs/protovalidate/schemas/adding-protovalidate/
[predefined-rules]: https://buf.build/docs/protovalidate/schemas/predefined-rules/
[standard-rules]: https://buf.build/docs/protovalidate/schemas/standard-rules/
[custom-rules]: https://buf.build/docs/protovalidate/schemas/custom-rules/
[connect-go]: https://buf.build/docs/protovalidate/quickstart/connect-go
[grpc-go]: https://buf.build/docs/protovalidate/quickstart/grpc-go
[grpc-java]: https://buf.build/docs/protovalidate/quickstart/grpc-java
[grpc-python]: https://buf.build/docs/protovalidate/quickstart/grpc-python
[bufstream]: https://buf.build/docs/protovalidate/quickstart/bufstream

[protoc-gen-validate]: https://github.com/bufbuild/protoc-gen-validate
[yup]: https://github.com/jquense/yup
[cel]: https://cel.dev
[go-validator]: https://github.com/go-playground/validator
[java-bean-validation]: https://beanvalidation.org/
[pydantic]: https://docs.pydantic.dev/

