package grpchandle

import (
	quantitygrpc "app/cmd/proxy-service/delivery/grpc/quantity"
	"app/generated/proto/servicegrpc"
	"app/internal/connection"
	constant "app/internal/constants"
	logapp "app/pkg/log"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Register() {
	connectInfo := connection.GetConnect()

	address := fmt.Sprintf("%s:%s", connectInfo.ProxyService.Host, connectInfo.ProxyService.Grpc)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logapp.Logger("listen-grpc-proxy", err.Error(), constant.ERROR_LOG)
		return
	}

	grpcServer := grpc.NewServer()
	servicegrpc.RegisterQuantityServiceServer(grpcServer, quantitygrpc.Register())

	log.Println("Start grpc proxy service: ", address)

	if err := grpcServer.Serve(lis); err != nil {
		logapp.Logger("start-grpc-proxy", err.Error(), constant.ERROR_LOG)
	}
}
