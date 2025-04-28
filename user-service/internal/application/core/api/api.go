package api

import (
	"context"

	"github.com/kida21/userservice/internal/application/core/domain"
	"github.com/kida21/userservice/internal/application/ports"
)


type Application struct {
	db ports.DBPort
}
func NewApplication(db ports.DBPort)(*Application){
   return &Application{db: db}
}

func (a *Application) Register(ctx context.Context,user domain.UserModel)( bool,error){
	registered,err:=a.db.Insert(ctx,&user)
	if err!=nil{
		return false,err
	}
	return registered,nil
}