package grpchandle

import (
	"app/generated/proto/servicegrpc"
	"app/internal/connection"
)

type grpcHandle struct {
	servicegrpc.UnimplementedStreamServiceServer
	infoConnection connection.Connection
}

func Register() servicegrpc.StreamServiceServer {
	return &grpcHandle{
		infoConnection: connection.GetConnect(),
	}
}
