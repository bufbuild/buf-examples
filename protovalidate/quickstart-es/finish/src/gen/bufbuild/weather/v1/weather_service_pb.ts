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

// @generated by protoc-gen-es v2.5.1 with parameter "target=ts,import_extension=js"
// @generated from file bufbuild/weather/v1/weather_service.proto (package bufbuild.weather.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv2";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv2";
import { file_buf_validate_validate } from "../../../buf/validate/validate_pb.js";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file bufbuild/weather/v1/weather_service.proto.
 */
export const file_bufbuild_weather_v1_weather_service: GenFile = /*@__PURE__*/
  fileDesc("CilidWZidWlsZC93ZWF0aGVyL3YxL3dlYXRoZXJfc2VydmljZS5wcm90bxITYnVmYnVpbGQud2VhdGhlci52MSKRAgoRR2V0V2VhdGhlclJlcXVlc3QSIQoIbGF0aXR1ZGUYASABKAJCD7pIDAoKHQAAtEItAAC0whIiCglsb25naXR1ZGUYAiABKAJCD7pIDAoKHQAANEMtAAA0wxK0AQoNZm9yZWNhc3RfZGF0ZRgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCgAG6SH26AXoKHWZvcmVjYXN0X2RhdGUud2l0aGluXzcyX2hvdXJzEitGb3JlY2FzdCBkYXRlIG11c3QgYmUgaW4gdGhlIG5leHQgNzIgaG91cnMuGix0aGlzID49IG5vdyAmJiB0aGlzIDw9IG5vdyArIGR1cmF0aW9uKCc3MmgnKWIGcHJvdG8z", [file_buf_validate_validate, file_google_protobuf_timestamp]);

/**
 * GetWeatherRequest is a request for weather at a point on Earth.
 *
 * @generated from message bufbuild.weather.v1.GetWeatherRequest
 */
export type GetWeatherRequest = Message<"bufbuild.weather.v1.GetWeatherRequest"> & {
  /**
   * latitude must be between -90 and 90, inclusive, to be valid. Use of a
   * float allows precision to about one square meter.
   *
   * @generated from field: float latitude = 1;
   */
  latitude: number;

  /**
   * longitude must be between -180 and 180, inclusive, to be valid. Use of a
   * float allows precision to about one square meter.
   *
   * @generated from field: float longitude = 2;
   */
  longitude: number;

  /**
   * forecast_date for the weather request. It must be within the next
   * three days.
   *
   * @generated from field: google.protobuf.Timestamp forecast_date = 3;
   */
  forecastDate?: Timestamp;
};

/**
 * Describes the message bufbuild.weather.v1.GetWeatherRequest.
 * Use `create(GetWeatherRequestSchema)` to create a new message.
 */
export const GetWeatherRequestSchema: GenMessage<GetWeatherRequest> = /*@__PURE__*/
  messageDesc(file_bufbuild_weather_v1_weather_service, 0);

