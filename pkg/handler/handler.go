package handler

import (
	v1 "exercise-go-api/pkg/handler/v1"
	"exercise-go-api/pkg/handler/version"
	"exercise-go-api/pkg/infrastracture/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	V1      *v1.Handler
	Version *version.Handler
}

func NewHandler(logger *zap.Logger, repo *persistence.Repositories, ver string) *Handler {
	h := &Handler{
		V1:      v1.NewHandler(logger.Named("v1"), repo),
		Version: version.NewHandler(logger.Named("version"), ver),
	}

	return h
}

// // DB なし
// func NewHandler(logger *zap.Logger) *Handler {
// 	h := &Handler{
// 		V1:      v1.NewHandler(logger.Named("v1")),
// 	}

// 	return h
// }