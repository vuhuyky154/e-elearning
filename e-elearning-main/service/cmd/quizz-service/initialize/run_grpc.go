package initialize

import (
	grpchandle "app/cmd/quizz-service/delivery/grpc"
	"app/generated/proto/servicegrpc"
	"app/internal/connection"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func runGRPC() {
	listenGRPC, err := net.Listen("tcp", fmt.Sprintf(":%s", connection.GetConnect().QuizzService.Grpc))
	if err != nil {
		log.Println("Error start quizz server grpc: ", err)
		return
	}

	grpcServer := grpc.NewServer()

	handleGrpc := grpchandle.Register()

	servicegrpc.RegisterQuizzServiceServer(grpcServer, handleGrpc)

	log.Println("Start quizz server grpc")
	grpcServer.Serve(listenGRPC)
}
