// Copyright 2020-2026 Buf Technologies, Inc.
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
	peoplev1 "github.com/bufbuild/buf-examples/protovalidate/rules-predefined/gen/bufbuild/people/v1"
)

func TestSuffixValidation(t *testing.T) {
	t.Parallel()

	expectations := map[string]struct {
		givenName  string
		middleName string
		familyName string
		title      string
		suffix     string
		valid      bool
	}{
		"Suffix at the limit is accepted": {
			givenName:  strings.Repeat("a", 50),
			middleName: OMIT,
			familyName: strings.Repeat("a", 50),
			title:      OMIT,
			suffix:     strings.Repeat("a", 40),
			valid:      true,
		},
		"An excessively long Suffix is rejected": {
			givenName:  strings.Repeat("a", 50),
			middleName: OMIT,
			familyName: strings.Repeat("a", 50),
			title:      OMIT,
			suffix:     strings.Repeat("a", 41),
			valid:      false,
		},
	}

	for testName, testCase := range expectations {
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			person := &peoplev1.Person{}
			person.SetGivenName(testCase.givenName)
			person.SetFamilyName(testCase.familyName)

			if testCase.middleName != OMIT {
				person.SetMiddleName(testCase.middleName)
			}

			if testCase.title != OMIT {
				person.SetTitle(testCase.title)
			}

			if testCase.suffix != OMIT {
				person.SetSuffix(testCase.suffix)
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
