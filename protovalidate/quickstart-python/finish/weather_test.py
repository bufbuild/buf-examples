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

import unittest

import bufbuild.weather.v1.weather_service_pb2 as weather_service_pb2

from datetime import datetime, timedelta
from protovalidate import ValidationError

from weather import validateWeather


class WeatherTest(unittest.TestCase):
    def test_valid_request(self):
        message = weather_service_pb2.GetWeatherRequest()
        message.latitude = 45
        message.longitude = 45
        message.forecast_date = datetime.now() + timedelta(days=1)

        try:
            validateWeather(message)
        except ValidationError:
            self.fail("validate() had a ValidationError")

    def test_bad_latitude(self):
        message = weather_service_pb2.GetWeatherRequest()
        message.latitude = -91
        message.longitude = 180
        message.forecast_date = datetime.now() + timedelta(days=1)

        with self.assertRaises(ValidationError) as context:
            validateWeather(message)

        validationErr = context.exception
        self.assertEqual(1, len(validationErr.violations))

    def test_bad_longitude(self):
        message = weather_service_pb2.GetWeatherRequest()
        message.latitude = -90
        message.longitude = 181
        message.forecast_date = datetime.now() + timedelta(days=1)

        with self.assertRaises(ValidationError) as context:
            validateWeather(message)

        validationErr = context.exception
        self.assertEqual(1, len(validationErr.violations))

    def test_bad_forecast_date(self):
        message = weather_service_pb2.GetWeatherRequest()
        message.latitude = -90
        message.longitude = 180
        message.forecast_date = datetime.now() + timedelta(days=4)

        with self.assertRaises(ValidationError) as context:
            validateWeather(message)

        validationErr = context.exception
        self.assertEqual(1, len(validationErr.violations))
