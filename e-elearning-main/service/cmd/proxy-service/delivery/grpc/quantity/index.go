package quantitygrpc

import (
	appcommon "app/cmd/proxy-service/app_common"
	"app/generated/proto/servicegrpc"
)

type server struct {
	servicegrpc.UnimplementedQuantityServiceServer
	listQuantityGrpc []servicegrpc.QuantityServiceClient
}

func Register() servicegrpc.QuantityServiceServer {
	return &server{
		listQuantityGrpc: appcommon.GetListGrpcQuantityGrpc(),
	}
}
