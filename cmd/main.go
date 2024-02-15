package main

import (
	"context"
	"fmt"
	"github.com/Murat993/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 5000

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	req.GetUsernames()
	return &chat_v1.CreateResponse{
		Id: 5,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort)) // tcp подключение по порту 50051
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                      // Создаем новый grpc сервер
	reflection.Register(s)                     // Включается параметр. где со стороны клиента можно посмотреть какие есть ендпоинты
	chat_v1.RegisterChatV1Server(s, &server{}) // Регистрируем сервер grpc

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // Serve - блокирует и скрип вечно вист с портом который указан в lis
		log.Fatalf("failed to serve: %v", err)
	}
}
