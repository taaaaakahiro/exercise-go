package repository

import (
	"context"
	
	"exercise-go-api/pkg/domain/entity"
)

type IUserRepository interface {
	ListUsers(ctx context.Context) ([]entity.User, error)
}