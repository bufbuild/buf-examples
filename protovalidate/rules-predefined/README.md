![The Buf logo](https://raw.githubusercontent.com/bufbuild/protovalidate/main/.github/buf-logo.svg)

# Using Protovalidate: Predefined rules

Predefined rules are complicated: they require proto2 or Edition 2023 but are often used in projects where most messages are proto3. 

We've created this directory to provide you with examples across Protobuf versions while also making sure the code samples
shown in Protovalidate's [predefined rules documentation][documentation] are correct. Every example here is tested in our CI/CD process and imported directly into Protovalidate.com.

## `additional-examples` directories

The nested `additional-examples` directory has additional examples with changes represented in a sparse directory tree which can, in turn, have its own `additional-examples` directory. Running `make ci` will perform validations on the current directory, and recursively copy itself and the rest of the boilerplate into the tree of additional-examples directories to construct and then run validations for each override scenario.

[documentation]: https://protovalidate.com/schemas/predefined-rules/

