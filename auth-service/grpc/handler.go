package grpc

import (
	"context"
	pb "github.com/kida21/authservice/grpc/gen"
	"github.com/kida21/authservice/internal/application/core/api"
	"github.com/kida21/authservice/internal/domain"
	
)
type Handler struct {
	service api.Service
	pb.UnimplementedAuthenticationServer
	
}
func NewHandler(service api.Service)(*Handler){
	return &Handler{service: service}
}
func (h *Handler) Login(ctx context.Context,req *pb.LoginRequest)(*pb.LoginResponse,error){
	tokenString,err:=h.service.AuthenticateUser(ctx,domain.Credential{Email: req.Email,Password: req.Password})
	return &pb.LoginResponse{Token: tokenString},err
}