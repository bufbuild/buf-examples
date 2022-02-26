# Protobuf linting

> You need to have the [`buf` CLI][install] installed to follow along with this example.

To lint the [`bad`](./bad) project:

```sh
buf lint ./bad
```

❌  Uh-oh! This module is still pretty rough:

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

To lint the [`good`](./good) project:

```sh
buf lint ./good
```

✅  Much better! This project has a clean bill of health.

[install]: https://docs.buf.build/installation
[lint]: https://docs.buf.build/lint
