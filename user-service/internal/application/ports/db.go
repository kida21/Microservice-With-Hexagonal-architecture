package ports

import (
	"context"
	"github.com/kida21/userservice/internal/application/core/domain"
)

type DBPort interface {
	Insert(ctx context.Context,user *domain.UserModel)(bool,error)
	Update(ctx context.Context,user *domain.UserModel)(*domain.UserModel,error)
	Delete(ctx context.Context,id int64)(bool,error)
}