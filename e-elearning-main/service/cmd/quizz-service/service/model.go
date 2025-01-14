package service

import (
	"app/internal/entity"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"
)

type Service struct {
	QueryQuizzService    query.QueryService[entity.Quizz]
	RawQueryQuizzService rawquery.QueryRawService[entity.Quizz]
}
