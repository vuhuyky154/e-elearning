package service

import (
	"app/internal/entity"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"
)

func Register() Service {
	return Service{
		QueryQuizzService:    query.Register[entity.Quizz](),
		RawQueryQuizzService: rawquery.Register[entity.Quizz](),
	}
}
