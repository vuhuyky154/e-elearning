package quantitygrpc

import (
	"app/generated/proto/servicegrpc"
	"context"
	"errors"
	"log"
)

func (s *server) InitProcessQuantity(ctx context.Context, req *servicegrpc.InitProcessQuantityRequest) (*servicegrpc.InitProcessQuantityResponse, error) {
	log.Println("proxy handle")
	if (len(s.listQuantityGrpc)) == 1 {
		return s.listQuantityGrpc[0].InitProcessQuantity(ctx, req)
	}

	return nil, errors.New("none service")
}
