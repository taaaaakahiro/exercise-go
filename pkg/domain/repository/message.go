package repository

import (
	"context"
	
	"exercise-go-api/pkg/domain/entity"
)

type IMessageRepository interface {
	ListMessages(ctx context.Context, userId int) ([]entity.Message, error)
}