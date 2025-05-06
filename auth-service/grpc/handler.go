package grpc

import (
	"context"
	//pb "github.com/kida21/authservice/grpc/gen"
	"github.com/kida21/authservice/internal/application/core/api"
	"github.com/kida21/authservice/internal/domain"
	authpb"github.com/kida21/Microservice-With-Hexagonal-architecture/z-proto/auth"
	
)
type Handler struct {
	service api.Service
	authpb.UnimplementedAuthenticationServer
	
}
func NewHandler(service api.Service)(*Handler){
	return &Handler{service: service}
}
func (h *Handler) Login(ctx context.Context,req *authpb.LoginRequest)(*authpb.LoginResponse,error){
	tokenString,err:=h.service.AuthenticateUser(ctx,domain.Credential{Email: req.Email,Password: req.Password})
	return &authpb.LoginResponse{Token: tokenString},err
}