package grpchandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	"app/generated/proto/sharedgrpc"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	logapp "app/pkg/log"

	"context"
)

func (s *grpcHandle) GetById(ctx context.Context, req *sharedgrpc.ID) (*servicegrpc.QuizzResponse, error) {
	quizz, err := s.service.QueryQuizzService.First(requestdata.QueryReq[entity.Quizz]{
		Condition: "id = ?",
		Args:      []interface{}{req.Id},
	})
	if err != nil {
		logapp.Logger("get-quizz-by-id", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	res := &servicegrpc.QuizzResponse{
		Id:         uint64(quizz.ID),
		Ask:        quizz.Ask,
		Time:       int32(quizz.Time),
		ResultType: enumgrpc.ResultType(enumgrpc.ResultType_value[string(quizz.ResultType)]),
		Result:     quizz.Result,
		Option:     quizz.Option,
		EntityType: enumgrpc.EntityType(enumgrpc.ResultType_value[string(quizz.EntityType)]),
		EntityId:   uint64(quizz.EntityId),
		CreatedAt:  quizz.CreatedAt.Unix(),
		UpdatedAt:  quizz.UpdatedAt.Unix(),
		DeletedAt:  quizz.DeletedAt.Time.Unix(),
	}

	return res, nil
}
