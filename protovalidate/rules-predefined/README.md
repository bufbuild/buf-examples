![The Buf logo](https://raw.githubusercontent.com/bufbuild/protovalidate/main/.github/buf-logo.svg)

# Using Protovalidate: Predefined rules

This directory contains companion code for [predefined rules documentation][documentation], where you can learn to refactor repeated Protovalidate rules into reusable predefined rules:

1. Create a rule file.
2. Extend a rule message to create a predefined rule.
3. Use your predefined rule.

[documentation]: https://protovalidate.com/schemas/predefined-rules/

## `supplementary` directories

The nested `supplementary` directory has additional examples with changes represented in a sparse directory tree which can, in turn, have its own `supplementary` directory. Running `make ci` will perform validations on the current directory, and recursively copy itself and the rest of the boilerplate into the tree of supplementary directories to construct and then run validations for each override scenario.
