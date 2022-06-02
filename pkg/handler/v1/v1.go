package v1

import (
	"exercise-go-api/pkg/infrastracture/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	repo   *persistence.Repositories
}

func NewHandler(logger *zap.Logger, repositories *persistence.Repositories) *Handler {
	return &Handler{
		logger: logger,
		repo:   repositories,
	}
}
// // DBなし
// func NewHandler(logger *zap.Logger) *Handler {
// 	return &Handler{
// 		logger: logger,
// 		// repo:   repositories,
// 	}
// }