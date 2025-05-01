package adapter

import (
	"context"
	"log"

	"github.com/kida21/authservice/internal/domain"
	userpb "github.com/kida21/userservice/gen"
	"google.golang.org/grpc"
)
type Adapter struct{
	conn *grpc.ClientConn
}
func NewAdapter(conn *grpc.ClientConn)(*Adapter){
	return &Adapter{conn: conn}
}
func (a *Adapter) Login(ctx context.Context,input *domain.Credential)(string,error){
    client:=userpb.NewUserClient(a.conn)
	response,err:=client.ValidateCredential(ctx,&userpb.ValidationRequest{Email: input.Email,Password: input.Password})
	if err!=nil{
		return "",err
	}
    if !response.Valid{
		log.Fatalf("Unauthorized")
	}
	//for debugging purpose
	return "token sent",nil
}