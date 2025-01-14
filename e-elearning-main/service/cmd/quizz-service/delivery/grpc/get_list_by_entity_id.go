package grpchandle

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	constant "app/internal/constants"
	requestdata "app/internal/dto/client"
	"app/internal/entity"
	logapp "app/pkg/log"
	"context"
)

func (h *grpcHandle) GetListByEntityId(ctx context.Context, req *servicegrpc.GetListQuizzRequest) (*servicegrpc.GetListByEntityIdResponse, error) {
	quizzs, err := h.service.QueryQuizzService.Find(requestdata.QueryReq[entity.Quizz]{
		Condition: "entity_id = ?",
		Args:      []interface{}{req.EntityId},
	})
	if err != nil {
		logapp.Logger("get-list-by-entity-id", err.Error(), constant.ERROR_LOG)
		return nil, err
	}

	listQuizz := []*servicegrpc.QuizzResponse{}
	for _, quizz := range quizzs {
		listQuizz = append(listQuizz, &servicegrpc.QuizzResponse{
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
		})
	}

	res := &servicegrpc.GetListByEntityIdResponse{
		Quizzs: listQuizz,
		Total:  int64(len(listQuizz)),
	}

	return res, nil
}
