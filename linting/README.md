# Protobuf linting

> You need to have the [`buf` CLI][install] installed to follow along with this example.

Buf enables you to [lint] Protobuf sources in accordance with [lint rules][rules] that you specify in your [`buf.yaml`][buf-yaml] configuration file. This project shows linting in action with two different [Buf inputs][inputs]:

* The [`bad`](./bad) input is a [module] that violates a whole host of lint rules. This module "works" but it ain't pretty.
* The [`good`](./good) module, on the other hand, violates no lint rules and is thus a much better citizen of the Protobuf universe.

Both inputs use the [`DEFAULT`][default] lint rules and are available on the [Buf Schema Registry][bsr]:

* [`buf.build/buf-examples/linting-bad`][bsr-bad]
* [`buf.build/buf-examples/linting-good`][bsr-good]

To lint the [`bad`](./bad) input:

```sh
buf lint ./bad
```

❌  Uh-oh! We can see that this module is pretty rough:

```
bad/acme/weather/v1/weather.proto:5:1:Files with package "weather" must be within a directory "weather" relative to root but were in directory "acme/weather/v1".
bad/acme/weather/v1/weather.proto:5:1:Package name "weather" should be suffixed with a correctly formed version, such as "weather.v1".
bad/acme/weather/v1/weather.proto:10:3:Enum value name "sunny" should be prefixed with "CONDITION_".
bad/acme/weather/v1/weather.proto:10:3:Enum value name "sunny" should be UPPER_SNAKE_CASE, such as "SUNNY".
bad/acme/weather/v1/weather.proto:10:3:Enum zero value name "sunny" should be suffixed with "_UNSPECIFIED".
bad/acme/weather/v1/weather.proto:11:3:Enum value name "rainy" should be prefixed with "CONDITION_".
bad/acme/weather/v1/weather.proto:11:3:Enum value name "rainy" should be UPPER_SNAKE_CASE, such as "RAINY".
bad/acme/weather/v1/weather.proto:12:3:Enum value name "other" should be prefixed with "CONDITION_".
bad/acme/weather/v1/weather.proto:12:3:Enum value name "other" should be UPPER_SNAKE_CASE, such as "OTHER".
bad/acme/weather/v1/weather.proto:40:9:Service name "WeatherVision" should be suffixed with "Service".
bad/acme/weather/v1/weather.proto:41:3:"weather.Weather" is used as the request or response type for multiple RPCs.
bad/acme/weather/v1/weather.proto:41:23:RPC request type "GetWeather" should be named "CurrentWeatherRequest" or "WeatherVisionCurrentWeatherRequest".
bad/acme/weather/v1/weather.proto:41:44:RPC response type "Weather" should be named "CurrentWeatherResponse" or "WeatherVisionCurrentWeatherResponse".
bad/acme/weather/v1/weather.proto:42:3:"weather.Weather" is used as the request or response type for multiple RPCs.
bad/acme/weather/v1/weather.proto:42:24:RPC request type "Expected" should be named "ExpectedWeatherRequest" or "WeatherVisionExpectedWeatherRequest".
bad/acme/weather/v1/weather.proto:42:43:RPC response type "Weather" should be named "ExpectedWeatherResponse" or "WeatherVisionExpectedWeatherResponse".
```

To lint the [`good`](./good) input:

```sh
buf lint ./good
```

✅  Much better! This project has a clean bill of health.

[bsr]: https://docs.buf.build/bsr
[bsr-bad]: https://buf.build/buf-examples/linting-bad
[bsr-good]: https://buf.build/buf-examples/linting-good
[buf-yaml]: https://docs.buf.build/configuration/v1/buf-yaml
[default]: https://docs.buf.build/lint/rules#default
[inputs]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
[jsonl]: https://jsonlines.org
[lint]: https://docs.buf.build/lint
[module]: https://docs.buf.build/bsr/overview#modules
[rules]: https://docs.buf.build/lint/rules
