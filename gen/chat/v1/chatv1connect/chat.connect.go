// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chat/v1/chat.proto

package chatv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/omihirofumi/grpc-connect-sample/gen/chat/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "chat.v1.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceSayProcedure is the fully-qualified name of the ChatService's Say RPC.
	ChatServiceSayProcedure = "/chat.v1.ChatService/Say"
)

// ChatServiceClient is a client for the chat.v1.ChatService service.
type ChatServiceClient interface {
	Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error)
}

// NewChatServiceClient constructs a client for the chat.v1.ChatService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		say: connect_go.NewClient[v1.SayRequest, v1.SayResponse](
			httpClient,
			baseURL+ChatServiceSayProcedure,
			opts...,
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	say *connect_go.Client[v1.SayRequest, v1.SayResponse]
}

// Say calls chat.v1.ChatService.Say.
func (c *chatServiceClient) Say(ctx context.Context, req *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error) {
	return c.say.CallUnary(ctx, req)
}

// ChatServiceHandler is an implementation of the chat.v1.ChatService service.
type ChatServiceHandler interface {
	Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error)
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	chatServiceSayHandler := connect_go.NewUnaryHandler(
		ChatServiceSayProcedure,
		svc.Say,
		opts...,
	)
	return "/chat.v1.ChatService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChatServiceSayProcedure:
			chatServiceSayHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("chat.v1.ChatService.Say is not implemented"))
}
