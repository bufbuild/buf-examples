import { expect, test } from "vitest";
import { create } from "@bufbuild/protobuf";
import {
  GetWeatherRequest,
  GetWeatherRequestSchema,
} from "./gen/bufbuild/weather/v1/weather_service_pb.js";
import { validateWeather } from "./weather.js";

test("validate method is defined", () => {
  const msg = create(GetWeatherRequestSchema);

  const result = validateWeather(msg);

  console.log(result);

  expect(result.kind).toEqual("valid");
});

test("another", () => {
  const msg = create(GetWeatherRequestSchema, {
    latitude: 95,
  });

  const result = validateWeather(msg);

  console.log(result);

  expect(result.kind).toEqual("invalid");
});
