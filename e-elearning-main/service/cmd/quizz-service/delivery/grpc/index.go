package grpchandle

import (
	"app/cmd/quizz-service/service"
	"app/generated/proto/servicegrpc"
)

type grpcHandle struct {
	servicegrpc.UnimplementedQuizzServiceServer
	service service.Service
}

func Register() servicegrpc.QuizzServiceServer {
	return &grpcHandle{
		service: service.Register(),
	}
}
