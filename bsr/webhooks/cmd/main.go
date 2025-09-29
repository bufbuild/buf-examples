package main

import (
	"context"
	"fmt"
	"net/http"

	webhookconnect "buf.build/gen/go/bufbuild/buf/connectrpc/go/buf/alpha/webhook/v1alpha1/webhookv1alpha1connect"
	registryv1alpha1 "buf.build/gen/go/bufbuild/buf/protocolbuffers/go/buf/alpha/registry/v1alpha1"
	webhookv1alpha1 "buf.build/gen/go/bufbuild/buf/protocolbuffers/go/buf/alpha/webhook/v1alpha1"
	connect "connectrpc.com/connect"
)

type webhookEventHandler struct{}

func (h *webhookEventHandler) Event(
	ctx context.Context,
	req *connect.Request[webhookv1alpha1.EventRequest],
) (*connect.Response[webhookv1alpha1.EventResponse], error) {
	// Handle the type-safe incoming request for the push event:
	payload := req.Msg
	switch payload.Event {
	case registryv1alpha1.WebhookEvent_WEBHOOK_EVENT_REPOSITORY_PUSH:
		pushEvent := payload.Payload.GetRepositoryPush()
		fmt.Println("received repo push event:", pushEvent)
	default:
		fmt.Println("unknown event:", payload.Event)
	}

	// Webhook listener has an empty response
	return connect.NewResponse(&webhookv1alpha1.EventResponse{}), nil
}

// Connect handler based on: https://connectrpc.com/docs/go/getting-started#implement-handler
func main() {
	mux := http.NewServeMux()
	mux.Handle(webhookconnect.NewEventServiceHandler(&webhookEventHandler{}))
	http.ListenAndServe("localhost:8080", mux)
}
