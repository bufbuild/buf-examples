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

package weather

import (
	"testing"
	"time"

	"buf.build/go/protovalidate"
	weatherv1 "github.com/bufbuild/buf-examples/protovalidate/quickstart-go/finish/gen/bufbuild/weather/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRequests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		request    *weatherv1.GetWeatherRequest
		violations []string
	}{{
		name: "valid",
		request: &weatherv1.GetWeatherRequest{
			Latitude:  45,
			Longitude: 45,
		},
	}, {
		name: "latitude too high",
		request: &weatherv1.GetWeatherRequest{
			Latitude:  91,
			Longitude: 45,
		},
		violations: []string{
			"value must be greater than or equal to -90 and less than or equal to 90",
		},
	}, {
		name: "latitude too low",
		request: &weatherv1.GetWeatherRequest{
			Latitude:  -180,
			Longitude: 45,
		},
		violations: []string{
			"value must be greater than or equal to -90 and less than or equal to 90",
		},
	}, {
		name: "too far in the future",
		request: &weatherv1.GetWeatherRequest{
			Latitude:     45,
			Longitude:    45,
			ForecastDate: timestamppb.New(time.Now().AddDate(0, 0, 4)),
		},
		violations: []string{
			"Forecast date must be in the next 72 hours.",
		},
	}}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			err := validateWeather(testCase.request)

			if len(testCase.violations) > 0 {
				require.Error(t, err)

				var validationError *protovalidate.ValidationError
				require.ErrorAs(t, err, &validationError)
				require.Len(t, validationError.Violations, len(testCase.violations))

				for i, violation := range testCase.violations {
					assert.Equal(t, violation, validationError.Violations[i].Proto.GetMessage())
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
