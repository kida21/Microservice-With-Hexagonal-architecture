package grpc

import (
	"github.com/kida21/authservice/internal/application/core/api"
    pb"github.com/kida21/authservice/grpc/gen"
)

type Handler struct {
	service api.Service
	pb.UnimplementedAuthenticationServer
}
func NewHandeler(service api.Service)(Handler){
	return Handler{service: service}
}