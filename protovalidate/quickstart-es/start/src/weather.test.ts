import { expect, test } from "vitest";
import { validateWeather } from "./weather.js";

test("validate method is defined", () => {
  expect(validateWeather).toBeDefined();
});
