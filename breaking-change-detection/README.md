# Breaking change detection

> You need to have the [`buf` CLI][install] installed to follow along with this example.

This project shows Buf [breaking change detection][breaking] in action. There are three different [Buf inputs][inputs] in play here:

* An [`initial`](./initial/acme/weather/v1/weather.proto) input that serves as the basis for comparison.
* An input called [`compatible`](./compatible/acme/weather/v1/weather.proto) that introduces non-breaking changes to `initial`. Inline comments indicate changes.
* An input called [`incompatible`](./incompatible/acme/weather/v1/weather.proto) that introduces breaking changes to `initial`. Inline comments indicate changes.

All three inputs are for the same `acme.weather.v1` package. Each input uses the default [`FILE`][file] breaking change rule, which detects breaking changes on a per-file basis and also verifies compatibility in how Protobuf definitions act across the wire (`WIRE`) and in the JSON representation of the definitions (`WIRE_JSON`).

To verify that the `compatible` input introduces no breaking changes against `initial`:

```sh
buf breaking ./compatible --against ./initial
```

✅  No output indicates that breaking change detection was successful.

To verify that `incompatible` introduces breaking changes against `initial`:

```sh
buf breaking ./incompatible --against ./initial
```

❌  Uh-oh! Breaking changes have indeed been introduced:

```
incompatible/acme/weather/v1/weather.proto:1:1:Previously present message "GetWeatherRequest" was deleted from file.
incompatible/acme/weather/v1/weather.proto:1:1:Previously present message "GetWeatherResponse" was deleted from file.
incompatible/acme/weather/v1/weather.proto:7:21:Enum value "1" on enum "Condition" changed name from "CONDITION_SUNNY" to "CONDITION_FOGGY".
incompatible/acme/weather/v1/weather.proto:12:3:Field "1" with name "latitude_min" on message "Location" changed option "json_name" from "latitude" to "latitudeMin".
incompatible/acme/weather/v1/weather.proto:12:9:Field "1" on message "Location" changed name from "latitude" to "latitude_min".
incompatible/acme/weather/v1/weather.proto:13:3:Field "2" with name "latitude_max" on message "Location" changed option "json_name" from "longitude" to "latitudeMax".
incompatible/acme/weather/v1/weather.proto:13:9:Field "2" on message "Location" changed name from "longitude" to "latitude_max".
incompatible/acme/weather/v1/weather.proto:18:1:Previously present reserved range "[5]" on message "Weather" was deleted.
incompatible/acme/weather/v1/weather.proto:36:1:Previously present RPC "GetWeather" on service "WeatherService" was deleted.
```

[breaking]: https://docs.buf.build
[file]: https://docs.buf.build/breaking/rules#categories
[inputs]: https://docs.buf.build/reference/inputs
[install]: https://docs.buf.build/installation
