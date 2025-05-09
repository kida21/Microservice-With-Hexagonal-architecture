package ports

import (
	"context"

	"github.com/kida21/authservice/internal/domain"
)

type AuthPort interface {
	Login(ctx context.Context,input *domain.Credential )(string,error)
}