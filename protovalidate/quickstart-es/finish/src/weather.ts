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

import {
  createValidator,
  type ValidationResult,
} from "@bufbuild/protovalidate";
import {
  GetWeatherRequest,
  GetWeatherRequestSchema,
} from "./gen/bufbuild/weather/v1/weather_service_pb.js";

const validator = createValidator();

export function validateWeather(msg: GetWeatherRequest): ValidationResult {
  return validator.validate(GetWeatherRequestSchema, msg);
}
