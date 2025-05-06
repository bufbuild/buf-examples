package main

import (
    "context"

    "buf.build/go/bufplugin/check"
    "buf.build/go/bufplugin/check/checkutil"
    "google.golang.org/protobuf/reflect/protoreflect"
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
    _ check.Request,
    methodDescriptor protoreflect.MethodDescriptor,
) error {
    responseWriter.AddAnnotation(
        check.WithMessage("hello world"),
    )
    return nil
}
