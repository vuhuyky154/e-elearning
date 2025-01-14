package grpchandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	"app/generated/proto/sharedgrpc"
	"app/internal/apperrors"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	logapp "app/pkg/log"
	"context"
)

func (h *grpcHandle) UpdateQuizz(ctx context.Context, req *servicegrpc.UpdateQuizzRequest) (*sharedgrpc.Empty, error) {
	if req.Payload.ResultType == enumgrpc.ResultType_QUIZZ_SINGLE_RESULT &&
		len(req.Payload.Result) != 1 {
		return nil, apperrors.ErrorSingleQuizzHaveMoreResult
	}

	_, err := h.service.QueryQuizzService.Update(requestdata.QueryReq[entity.Quizz]{
		Data: entity.Quizz{
			Ask:        req.Payload.Ask,
			ResultType: entity.RESULT_TYPE(req.Payload.ResultType.String()),
			Result:     req.Payload.Result,
			Option:     req.Payload.Option,
			Time:       int(req.Payload.Time),
		},
		Condition: "id = ?",
		Args:      []interface{}{req.Id},
	})
	if err != nil {
		logapp.Logger("update-quizz", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	return &sharedgrpc.Empty{}, nil
}
