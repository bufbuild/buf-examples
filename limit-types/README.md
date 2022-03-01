# Limit types

> You need to install the [`buf` CLI][install] to follow along with this example. Requires v1.1.0 or
> above.

The [`buf build`][build] command creates a [Buf image][image] (or
[`FileDescriptorSet`][filedescriptorset]) out of a [Buf input][input]. By default, the image or
`FileDescriptorSet` includes _all_ of the Protobuf types—messages, enums, and services—in the input.
But the `buf build` command offers [`--type`][type] option that enables you to supply fully
qualified Protobuf type names and limit the output to only the Protobuf descriptors required to
represent those types and their required dependencies.

In this example, we can use `buf build --type` to view different representations (as JSON) for
various Protobuf types in the [`auth/v1`](./auth/v1) module.

The `auth/v1` module has a [`User`](./auth/v1#L5-9) message with no dependencies. To display a JSON
representation of this type:

```sh
buf build \
  --type auth.v1.User \
  --output -#format=json | \
  jq '.file[0].messageType[0]'
```

The output:

```javascript
{
  "name": "User",
  "field": [
    {
      "name": "user_id",
      "number": 1,
      "label": "LABEL_OPTIONAL", // All fields are implicitly optional in proto3
      "type": "TYPE_STRING",
      "jsonName": "userId",
      "options": {
        "deprecated": true
      }
    },
    {
      "name": "username",
      "number": 2,
      "label": "LABEL_OPTIONAL",
      "type": "TYPE_STRING",
      "jsonName": "username"
    },
    {
      "name": "email",
      "number": 3,
      "label": "LABEL_OPTIONAL",
      "type": "TYPE_STRING",
      "jsonName": "email"
    }
  ]
}
```

Now, let's look at the output for the [`AuthenticateResponse`](./auth/v1#L18-30) message, which
has dependencies on other types in the module:

```sh
buf build \
  --type auth.v1.AuthenticateResponse \
  --output -#format=json | \
  jq '.file[0]'
```

We can see that the output is more involved, including type definitions for the parent
`AuthenticateResponse` message but also the `User` message, as above, and a `Result` enum:
above, but also 

```json
{
  "name": "auth/v1/auth.proto",
  "package": "auth.v1",
  "messageType": [
    {
      "name": "User",
      "field": [
        {
          "name": "user_id",
          "number": 1,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "userId",
          "options": {
            "deprecated": true
          }
        },
        {
          "name": "username",
          "number": 2,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "username"
        },
        {
          "name": "email",
          "number": 3,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "email"
        }
      ]
    },
    {
      "name": "AuthenticateResponse",
      "field": [
        {
          "name": "result",
          "number": 1,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_ENUM",
          "typeName": ".auth.v1.AuthenticateResponse.Result",
          "jsonName": "result"
        },
        {
          "name": "user",
          "number": 2,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_MESSAGE",
          "typeName": ".auth.v1.User",
          "jsonName": "user"
        }
      ],
      "enumType": [
        {
          "name": "Result",
          "value": [
            {
              "name": "RESULT_UNSPECIFIED",
              "number": 0
            },
            {
              "name": "RESULT_AUTHENTICATION_SUCCESS",
              "number": 1
            },
            {
              "name": "RESULT_USER_NOT_FOUND",
              "number": 2
            },
            {
              "name": "RESULT_MALFORMED_REQUEST",
              "number": 3
            },
            {
              "name": "RESULT_AUTHENTICATION_FAILED",
              "number": 4
            },
            {
              "name": "RESULT_INCORRECT_PASSWORD",
              "number": 5,
              "options": {
                "deprecated": true
              }
            }
          ]
        }
      ]
    }
  ],
  "syntax": "proto3",
  "bufExtension": {
    "isImport": false,
    "isSyntaxUnspecified": false
  }
}
```

Finally, let's look at the output for the [`AuthenticationService`](./auth/v1#L32-34) service, which
uses all of the available child types in the module:

```sh
buf build --type auth.v1.AuthenticationService --output -#format=json | jq '.file[0]'
```

The output:

```json
{
  "name": "auth/v1/auth.proto",
  "package": "auth.v1",
  "messageType": [
    {
      "name": "User",
      "field": [
        {
          "name": "user_id",
          "number": 1,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "userId",
          "options": {
            "deprecated": true
          }
        },
        {
          "name": "username",
          "number": 2,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "username"
        },
        {
          "name": "email",
          "number": 3,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "email"
        }
      ]
    },
    {
      "name": "AuthenticateRequest",
      "field": [
        {
          "name": "user_id",
          "number": 1,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "userId",
          "options": {
            "deprecated": true
          }
        },
        {
          "name": "username",
          "number": 2,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "username"
        },
        {
          "name": "email",
          "number": 3,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "email"
        },
        {
          "name": "password",
          "number": 4,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_STRING",
          "jsonName": "password"
        }
      ]
    },
    {
      "name": "AuthenticateResponse",
      "field": [
        {
          "name": "result",
          "number": 1,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_ENUM",
          "typeName": ".auth.v1.AuthenticateResponse.Result",
          "jsonName": "result"
        },
        {
          "name": "user",
          "number": 2,
          "label": "LABEL_OPTIONAL",
          "type": "TYPE_MESSAGE",
          "typeName": ".auth.v1.User",
          "jsonName": "user"
        }
      ],
      "enumType": [
        {
          "name": "Result",
          "value": [
            {
              "name": "RESULT_UNSPECIFIED",
              "number": 0
            },
            {
              "name": "RESULT_AUTHENTICATION_SUCCESS",
              "number": 1
            },
            {
              "name": "RESULT_USER_NOT_FOUND",
              "number": 2
            },
            {
              "name": "RESULT_MALFORMED_REQUEST",
              "number": 3
            },
            {
              "name": "RESULT_AUTHENTICATION_FAILED",
              "number": 4
            },
            {
              "name": "RESULT_INCORRECT_PASSWORD",
              "number": 5,
              "options": {
                "deprecated": true
              }
            }
          ]
        }
      ]
    }
  ],
  "service": [
    {
      "name": "AuthenticationService",
      "method": [
        {
          "name": "Authenticate",
          "inputType": ".auth.v1.AuthenticateRequest",
          "outputType": ".auth.v1.AuthenticateResponse",
          "options": {}
        }
      ]
    }
  ],
  "syntax": "proto3",
  "bufExtension": {
    "isImport": false,
    "isSyntaxUnspecified": false
  }
}
```

[build]: https://docs.buf.build/build/usage
[filedescriptorset]: https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto
[image]: https://docs.buf.build/reference/images
[input]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
[type]: https://docs.buf.build/build/usage#limit-to-specific-types
