package grpchandle

import (
	appcommon "app/cmd/stream-service/app_common"
	"app/generated/proto/servicegrpc"
	constant "app/internal/constants"
	logapp "app/pkg/log"
	"context"
	"fmt"
)

func (h *grpcHandle) InitProcess(ctx context.Context, req *servicegrpc.InitProcessRequest) (*servicegrpc.InitProcessResponse, error) {
	err := appcommon.CreateProcess(req.Uuid)
	if err != nil {
		logapp.Logger("goroutine-process-stream", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	ipServer := fmt.Sprintf(
		"%s:%s",
		h.infoConnection.StreamService.Host,
		h.infoConnection.StreamService.Grpc,
	)

	res := &servicegrpc.InitProcessResponse{
		Ip: ipServer,
	}

	return res, nil
}
