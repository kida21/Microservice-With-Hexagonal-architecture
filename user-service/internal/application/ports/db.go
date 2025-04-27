package ports

import "github.com/kida21/userservice/internal/application/core/domain"

type DBPort interface {
	Insert(*domain.UserModel)(domain.UserModel,error)
}