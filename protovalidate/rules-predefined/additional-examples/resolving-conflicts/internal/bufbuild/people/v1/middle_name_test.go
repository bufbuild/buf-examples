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
	"testing"

	"buf.build/go/protovalidate"
	peoplev1 "github.com/bufbuild/buf-examples/protovalidate/rules-predefined/gen/bufbuild/people/v1"
)

func TestMiddleNameValidation(t *testing.T) {
	t.Parallel()

	t.Run("Empty MiddleName is permitted", func(t *testing.T) {
		t.Parallel()

		person := &peoplev1.Person{}
		person.SetGivenName("George")
		person.SetMiddleName("")
		person.SetFamilyName("Washington")

		err := protovalidate.Validate(person)
		if err != nil {
			t.Fatalf("unexpected validation error received: %v", err)
		}
	})
}
