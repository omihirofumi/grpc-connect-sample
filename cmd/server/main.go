package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	chatv1 "github.com/omihirofumi/grpc-connect-sample/gen/chat/v1"
	"github.com/omihirofumi/grpc-connect-sample/gen/chat/v1/chatv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type ChatServer struct{}

func (s *ChatServer) Say(ctx context.Context, req *connect.Request[chatv1.SayRequest]) (*connect.Response[chatv1.SayResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&chatv1.SayResponse{
		Sentence: fmt.Sprintf("I said %s", req.Msg.Sentence),
	})
	res.Header().Set("Chat-Version", "v1")
	return res, nil
}

func main() {
	s := &ChatServer{}
	mux := http.NewServeMux()
	path, handler := chatv1connect.NewChatServiceHandler(s)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
