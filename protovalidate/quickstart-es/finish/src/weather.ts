import {
  GetWeatherRequest,
  GetWeatherRequestSchema,
} from "./gen/bufbuild/weather/v1/weather_service_pb.js";
import { createValidator } from "@bufbuild/protovalidate";

const validator = createValidator();

export function validateWeather(msg: GetWeatherRequest) {
  return validator.validate(GetWeatherRequestSchema, msg);
}
