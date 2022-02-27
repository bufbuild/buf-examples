# Formatting Protobuf sources

> You need to have the [`buf` CLI][install] installed to follow along with this example.

Buf enables you to [format] Protobuf sources using the `buf format` command. The formatter is _not_ configurable and uses only an opinionated set of rules.

## Diffs

To output a diff:

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

To fix files in place:

```sh
buf format --write ./bad
```

[format]: https://docs.buf.build/format
[install]: https://docs.buf.build/installation
