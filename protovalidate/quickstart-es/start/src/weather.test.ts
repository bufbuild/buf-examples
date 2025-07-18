// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { expect, test } from "vitest";
import { create } from "@bufbuild/protobuf";
import { timestampFromDate } from "@bufbuild/protobuf/wkt";
import { GetWeatherRequestSchema } from "./gen/bufbuild/weather/v1/weather_service_pb.js";
import { validateWeather } from "./weather.js";

test("valid message", () => {
  const msg = create(GetWeatherRequestSchema, {
    latitude: 45,
    longitude: 45,
  });

  const result = validateWeather(msg);

  expect(result).toBeDefined();
  expect(result.kind).toEqual("valid");
  expect(result.violations).toBeUndefined();
});

test("latitude too high", () => {
  const msg = create(GetWeatherRequestSchema, {
    latitude: 91,
    longitude: 45,
  });

  const result = validateWeather(msg);

  expect(result).toBeDefined();
  expect(result.kind).toEqual("invalid");

  const violations = result.violations ?? [];
  expect(violations.length).toEqual(1);
  expect(violations[0].message).toEqual(
    "value must be greater than or equal to -90 and less than or equal to 90",
  );
});

test("latitude too low", () => {
  const msg = create(GetWeatherRequestSchema, {
    latitude: -180,
    longitude: 45,
  });

  const result = validateWeather(msg);

  expect(result).toBeDefined();
  expect(result.kind).toEqual("invalid");

  const violations = result.violations ?? [];
  expect(violations.length).toEqual(1);
  expect(violations[0].message).toEqual(
    "value must be greater than or equal to -90 and less than or equal to 90",
  );
});

test("too far in the future", () => {
  // Create a date 7 days in the future
  const dt = new Date();
  dt.setDate(dt.getDate() + 7);

  const msg = create(GetWeatherRequestSchema, {
    latitude: 45,
    longitude: 45,
    forecastDate: timestampFromDate(dt),
  });

  const result = validateWeather(msg);

  expect(result).toBeDefined();
  expect(result.kind).toEqual("invalid");

  const violations = result.violations ?? [];
  expect(violations.length).toEqual(1);
  expect(violations[0].message).toEqual(
    "Forecast date must be in the next 72 hours.",
  );
});
