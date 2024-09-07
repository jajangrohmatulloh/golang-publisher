package repository

import (
	"publisher/model/domain"
	"context"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (domain.User, error)
}
