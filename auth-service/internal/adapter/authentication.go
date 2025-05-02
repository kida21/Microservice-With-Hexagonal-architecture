package adapter

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
type Claims struct {
	userId int64
	jwt.RegisteredClaims
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
	expirationTime:=time.Now().Add(time.Minute * 10)
   claim:=&Claims{userId: response.UserId,
	              RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)},
				  }
     key:=[]byte(os.Getenv("KEY"))
	 token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	 tokenString,err:=token.SignedString(key)
	 if err!=nil{
		log.Fatal(err)
	 }
	return tokenString,nil
}