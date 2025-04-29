package grpc

import (
	"context"
	pb "github.com/kida21/userservice/internal/adapter/grpc/gen"
	"github.com/kida21/userservice/internal/application/core/api"
	"github.com/kida21/userservice/internal/application/core/domain"
)
type Handler struct {
  service *api.Application
 pb.UnimplementedUserServer
}

func NewHandler(service *api.Application)(*Handler){
	return &Handler{service: service}
}
 
func (h *Handler) RegisterUser(ctx context.Context,req *pb.RegisterRequest)(*pb.RegisterResponse,error){
	user:=domain.UserModel{FirstName:req.Firstname,LastName: req.Lastname,Email: req.Email}
	if err:=user.Password.Set(req.Password);err!=nil{
		return &pb.RegisterResponse{},err
	}
    created,err:=h.service.RegisterUser(ctx,&user)
	if err!=nil{
		return &pb.RegisterResponse{},err
	}
	return &pb.RegisterResponse{Created:created},nil
}
func (h *Handler) ValidateCredential(ctx context.Context,req *pb.ValidationRequest)(*pb.ValidationResponse,error){
	input:=&domain.UserCredential{
		Email: req.Email,
		Password: req.Password,
	}
	 userId,valid,err:=h.service.ValidateUser(ctx,input)
	 if err!=nil{
		return &pb.ValidationResponse{},err
	 }
	 return &pb.ValidationResponse{UserId:userId,Valid: valid},nil
}