package api

import (
	"context"

	"github.com/kida21/authservice/internal/domain"
	"github.com/kida21/authservice/internal/ports"
)

type Service struct {
	auth ports.AuthPort
}

func NewService(auth ports.AuthPort)(Service){
	return Service{auth: auth}
}

func (s *Service) AuthenticateUser(ctx context.Context,input domain.Credential)()