package grpchandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	"app/generated/proto/sharedgrpc"
	"app/internal/apperrors"
	constant "app/internal/constants"
	"app/internal/entity"
	logapp "app/pkg/log"
	"context"
)

func (h *grpcHandle) CreateQuizz(ctx context.Context, req *servicegrpc.CreateQuizzRequest) (*sharedgrpc.ID, error) {
	if req.ResultType == enumgrpc.ResultType_QUIZZ_SINGLE_RESULT &&
		len(req.Result) != 1 {
		return nil, apperrors.ErrorSingleQuizzHaveMoreResult
	}

	newQuizz, err := h.service.QueryQuizzService.Create(entity.Quizz{
		Ask:        req.Ask,
		Time:       int(req.Time),
		Option:     req.Option,
		ResultType: entity.RESULT_TYPE(req.ResultType.String()),
		Result:     req.Result,
		EntityType: entity.ENTITY_TYPE(req.EntityType.String()),
		EntityId:   uint(req.EntityId),
	})

	if err != nil {
		logapp.Logger("create-quizz", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	return &sharedgrpc.ID{
		Id: uint64(newQuizz.ID),
	}, nil
}
