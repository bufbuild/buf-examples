# Formatting Protobuf sources

> You need to have the [`buf` CLI][install] installed to follow along with this example.

The `buf` CLI enables you to format Protobuf sources using the `buf format` command. It works just like formatters in other languages, such as [gofmt] for Go and [Prettier] for JavaScript. Please be aware, though, that the formatter is _not_ configurable and applies only an opinionated set of built-in rules.

## Creating diffs

With `buf format`, you can output diffs that display which changes the formatter suggests. To output a diff:

```sh
buf format --diff ./bad
```

```diff
diff -u bad/acme/weather/v1/weather.proto.orig bad/acme/weather/v1/weather.proto
--- bad/acme/weather/v1/weather.proto.orig	2022-02-26 16:48:17.000000000 -0800
+++ bad/acme/weather/v1/weather.proto	2022-02-26 16:48:17.000000000 -0800
@@ -1,36 +1,41 @@
 syntax = "proto3";
-import "google/protobuf/timestamp.proto";
+
 package acme.weather.v1;
 
+import "google/protobuf/timestamp.proto";
+
 message Location {
   float latitude = 1;
   float longitude = 2;
 }
+
 message CurrentWeatherRequest {
-    Location location = 1; }
+  Location location = 1;
+}
+
 message CurrentWeatherResponse {
   float temperature = 1;
   Condition condition = 2;
 }
+
 message ExpectedWeatherResponse {
   float temperate = 1;
-  Condition condition = 2; }
+  Condition condition = 2;
+}
+
 message ExpectedWeatherRequest {
   Location location = 1;
   google.protobuf.Timestamp time = 2;
 }
+
 enum Condition {
   CONDITION_UNSPECIFIED = 0;
   CONDITION_SUNNY = 1;
-    CONDITION_RAINY = 2;
+  CONDITION_RAINY = 2;
   CONDITION_OTHER = 3;
 }
 
-
-
 service WeatherVisionService {
-  rpc CurrentWeather (CurrentWeatherRequest)
-    returns (CurrentWeatherResponse);
-  rpc ExpectedWeather (ExpectedWeatherRequest)
-    returns (ExpectedWeatherResponse);
+  rpc CurrentWeather(CurrentWeatherRequest) returns (CurrentWeatherResponse);
+  rpc ExpectedWeather(ExpectedWeatherRequest) returns (ExpectedWeatherResponse);
 }
```

## Fixing files in place

With `buf format`, you can also change Protobuf files in place:

```sh
buf format --write ./bad
```

This rewrites all of the `.proto` files in the target input.

[gofmt]: https://pkg.go.dev/cmd/gofmt
[install]: https://docs.buf.build/installation
[prettier]: https://prettier.io
