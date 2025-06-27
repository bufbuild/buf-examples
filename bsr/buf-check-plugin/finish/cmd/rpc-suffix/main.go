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

package main

import (
	"context"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/option"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	defaultForbiddenRPCSuffix     = "Method"
	forbiddenRPCSuffixesOptionKey = "forbidden_rpc_suffixes"
)

var (
	rpcSuffixRuleSpec = &check.RuleSpec{
		ID:      "RPC_SUFFIX",
		Default: true,
		Purpose: "Checks that RPC names don't end with an illegal suffix.",
		Type:    check.RuleTypeLint,
		Handler: checkutil.NewMethodRuleHandler(checkRPCSuffix, checkutil.WithoutImports()),
	}
)

func main() {
	check.Main(&check.Spec{
		Rules: []*check.RuleSpec{
			rpcSuffixRuleSpec,
		},
	})
}

func checkRPCSuffix(
	_ context.Context,
	responseWriter check.ResponseWriter,
	request check.Request,
	methodDescriptor protoreflect.MethodDescriptor,
) error {
	methodName := string(methodDescriptor.Name())
	forbiddenRPCSuffixes, err := option.GetStringSliceValue(request.Options(), forbiddenRPCSuffixesOptionKey)
	if err != nil {
		return err
	}
	if len(forbiddenRPCSuffixes) == 0 {
		forbiddenRPCSuffixes = append(forbiddenRPCSuffixes, defaultForbiddenRPCSuffix)
	}
	for _, forbiddenRPCSuffix := range forbiddenRPCSuffixes {
		if strings.HasSuffix(methodName, forbiddenRPCSuffix) {
			responseWriter.AddAnnotation(
				check.WithDescriptor(methodDescriptor),
				check.WithMessagef("method name should not end with %q", forbiddenRPCSuffix),
			)
		}
	}
	return nil
}
