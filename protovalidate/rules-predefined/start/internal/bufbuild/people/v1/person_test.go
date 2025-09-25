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

package peoplev1

import (
	"strings"
	"testing"

	"buf.build/go/protovalidate"
	peoplev1 "github.com/bufbuild/buf-examples/protovalidate/rules-predefined/start/gen/bufbuild/people/v1"
)

func TestPersonValidation(t *testing.T) {
	t.Parallel()

	expectations := map[string]struct {
		firstName  string
		lastName   string
		middleName string
		title      string
		valid      bool
	}{
		"Empty FirstName is rejected": {
			firstName: "",
			lastName:  strings.Repeat("a", 50),
			title:     strings.Repeat("a", 50),
			valid:     false,
		},
		"An excessively long FirstName is rejected": {
			firstName: strings.Repeat("a", 51),
			lastName:  strings.Repeat("a", 50),
			title:     strings.Repeat("a", 50),
			valid:     false,
		},
		"FirstName at the limit is accepted": {
			firstName: strings.Repeat("a", 50),
			lastName:  strings.Repeat("a", 50),
			title:     strings.Repeat("a", 50),
			valid:     true,
		},
		"FirstName string under the limit is accepted": {
			firstName: strings.Repeat("a", 49),
			lastName:  strings.Repeat("a", 50),
			title:     strings.Repeat("a", 50),
			valid:     true,
		},
		"MiddleName under the limit is accepted": {
			firstName:  strings.Repeat("a", 49),
			lastName:   strings.Repeat("a", 50),
			middleName: strings.Repeat("a", 49),
			title:      strings.Repeat("a", 50),
			valid:      true,
		},
		"An excessively long MiddleName is rejected": {
			firstName:  strings.Repeat("a", 50),
			lastName:   strings.Repeat("a", 50),
			middleName: strings.Repeat("a", 51),
			title:      strings.Repeat("a", 50),
			valid:      false,
		},
		"An excessively long title is rejected": {
			firstName: strings.Repeat("a", 50),
			lastName:  strings.Repeat("a", 50),
			title:     strings.Repeat("a", 65),
			valid:     false,
		},
	}

	for testName, testCase := range expectations {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			person := &peoplev1.Person{
				FirstName:  testCase.firstName,
				LastName:   testCase.lastName,
				MiddleName: testCase.middleName,
				Title:      testCase.title,
			}

			err := protovalidate.Validate(person)

			if err != nil && testCase.valid {
				t.Fatalf("unexpected validation error received: %v", err)
			}

			if err == nil && !testCase.valid {
				t.Fatal("expected validation error")
			}
		})
	}
}
