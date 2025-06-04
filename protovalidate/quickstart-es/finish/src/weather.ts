import { createValidator, ValidationResult } from "@bufbuild/protovalidate";
import {
  GetWeatherRequest,
  GetWeatherRequestSchema,
} from "./gen/bufbuild/weather/v1/weather_service_pb.js";

const validator = createValidator();

export function validateWeather(msg: GetWeatherRequest): ValidationResult {
  return validator.validate(GetWeatherRequestSchema, msg);
}
