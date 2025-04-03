// Copyright 2025 Buf Technologies, Inc.
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

package com.bufbuild.weather;

import static org.junit.jupiter.api.Assertions.*;

import build.buf.protovalidate.ValidationResult;
import com.bufbuild.weather.v1.GetWeatherRequest;
import com.google.protobuf.Timestamp;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import org.junit.jupiter.api.Test;

public class WeatherTest {
    @Test
    public void TestValidRequest() {
        GetWeatherRequest message = GetWeatherRequest.newBuilder()
                .setLatitude(45)
                .setLongitude(45)
                .setForecastDate(Timestamp.newBuilder()
                        .setSeconds(Instant.now().plus(1, ChronoUnit.DAYS).getEpochSecond())
                        .build())
                .build();

        ValidationResult validationResult =
                assertDoesNotThrow(() -> new WeatherService().validateGetWeatherRequest(message));

        assertTrue(validationResult.isSuccess());
    }

    @Test
    public void TestBadLatitude() {
        GetWeatherRequest message = GetWeatherRequest.newBuilder()
                .setLatitude(-91)
                .setLongitude(45)
                .setForecastDate(Timestamp.newBuilder()
                        .setSeconds(Instant.now().plus(1, ChronoUnit.DAYS).getEpochSecond())
                        .build())
                .build();

        ValidationResult validationResult =
                assertDoesNotThrow(() -> new WeatherService().validateGetWeatherRequest(message));

        assertFalse(validationResult.isSuccess());
        assertEquals(1, validationResult.getViolations().size());
    }

    @Test
    public void TestBadLongitude() {
        GetWeatherRequest message = GetWeatherRequest.newBuilder()
                .setLatitude(-90)
                .setLongitude(181)
                .setForecastDate(Timestamp.newBuilder()
                        .setSeconds(Instant.now().plus(1, ChronoUnit.DAYS).getEpochSecond())
                        .build())
                .build();

        ValidationResult validationResult =
                assertDoesNotThrow(() -> new WeatherService().validateGetWeatherRequest(message));

        assertFalse(validationResult.isSuccess());
        assertEquals(1, validationResult.getViolations().size());
    }

    @Test
    public void TestBadForecastDate() {
        GetWeatherRequest message = GetWeatherRequest.newBuilder()
                .setLatitude(45)
                .setLongitude(45)
                .setForecastDate(Timestamp.newBuilder()
                        .setSeconds(Instant.now().plus(-1, ChronoUnit.DAYS).getEpochSecond())
                        .build())
                .build();

        ValidationResult validationResult =
                assertDoesNotThrow(() -> new WeatherService().validateGetWeatherRequest(message));

        assertFalse(validationResult.isSuccess());
        assertEquals(1, validationResult.getViolations().size());
    }
}
