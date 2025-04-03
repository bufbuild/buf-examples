# Copyright 2020-2025 Buf Technologies, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: bufbuild/weather/v1/weather_service.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'bufbuild/weather/v1/weather_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)bufbuild/weather/v1/weather_service.proto\x12\x13\x62ufbuild.weather.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8e\x01\n\x11GetWeatherRequest\x12\x1a\n\x08latitude\x18\x01 \x01(\x02R\x08latitude\x12\x1c\n\tlongitude\x18\x02 \x01(\x02R\tlongitude\x12?\n\rforecast_date\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0c\x66orecastDateb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'bufbuild.weather.v1.weather_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_GETWEATHERREQUEST']._serialized_start=100
  _globals['_GETWEATHERREQUEST']._serialized_end=242
# @@protoc_insertion_point(module_scope)
