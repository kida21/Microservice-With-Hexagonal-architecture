package ports

import "github.com/kida21/userservice/internal/application/core/domain"

type ApiPort interface {
	Register(domain.UserModel)(domain.UserModel,error)
}