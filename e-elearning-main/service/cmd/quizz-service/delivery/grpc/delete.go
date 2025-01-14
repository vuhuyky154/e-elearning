package grpchandle

import (
	"app/generated/proto/sharedgrpc"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	logapp "app/pkg/log"
	"context"
)

func (h *grpcHandle) DeleteById(ctx context.Context, req *sharedgrpc.ID) (*sharedgrpc.Empty, error) {
	err := h.service.QueryQuizzService.Delete(requestdata.QueryReq[entity.Quizz]{
		Condition: "id = ?",
		Args:      []interface{}{req.Id},
	})
	if err != nil {
		logapp.Logger("delete-quizz-by-id", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	return nil, nil
}
