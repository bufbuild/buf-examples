package main

import (
	"context"
	"fmt"
	"net/http"

	webhookconnect "buf.build/gen/go/bufbuild/buf/connectrpc/gosimple/buf/alpha/webhook/v1alpha1/webhookv1alpha1connect"
	registryv1alpha1 "buf.build/gen/go/bufbuild/buf/protocolbuffers/go/buf/alpha/registry/v1alpha1"
	webhookv1alpha1 "buf.build/gen/go/bufbuild/buf/protocolbuffers/go/buf/alpha/webhook/v1alpha1"
)

type webhookEventHandler struct{}

func (h *webhookEventHandler) Event(
	_ context.Context,
	req *webhookv1alpha1.EventRequest,
) (*webhookv1alpha1.EventResponse, error) {
	// Handle the type-safe incoming request for the push event:
	switch req.GetEvent() {
	case registryv1alpha1.WebhookEvent_WEBHOOK_EVENT_REPOSITORY_PUSH:
		pushEvent := req.GetPayload().GetRepositoryPush()
		fmt.Println("received repo push event:", pushEvent)
	default:
		fmt.Println("unknown event:", req.GetEvent())
	}

	// Webhook listener has an empty response
	return &webhookv1alpha1.EventResponse{}, nil
}

// Connect handler based on: https://connectrpc.com/docs/go/getting-started#implement-handler
func main() {
	mux := http.NewServeMux()
	mux.Handle(webhookconnect.NewEventServiceHandler(&webhookEventHandler{}))
	http.ListenAndServe("localhost:8080", mux)
}
