package main

import (
	"context"
	"log"
	"net"

	"github.com/serpis1/go_cook_book/formats/grpc/oneWay/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error with listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, ChatServer{})
	_ = grpcServer.Serve(lis)
}

type ChatServer struct{}

func (c ChatServer) SendMessage(ctx context.Context, msg *chat.ChatMesage) (*chat.ChatMesage, error) {
	return &chat.ChatMesage{
		Id:   1,
		Text: msg.Text,
	}, nil
}
