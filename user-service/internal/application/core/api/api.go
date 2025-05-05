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

func (a *Application) RegisterUser(ctx context.Context,user *domain.UserModel)( bool,error){
	registered,err:=a.db.Insert(ctx,user)
	if err!=nil{
		return false,err
	}
	return registered,nil
}
func (a *Application) UpdateUser(ctx context.Context,user *domain.UserModel)(int64,int64,error){
	id,version,err:=a.db.Update(ctx,user)
	if err!=nil{
		return 0,0,err
	}
	return id,version,nil
}
func(a *Application) DeleteUser(ctx context.Context,id int64)(bool,error){
	Deleted,err:=a.db.Delete(ctx,id)
	if err!=nil{
		return false,err
	}
	return Deleted,nil
}
func (a *Application) ValidateUser(ctx context.Context,input *domain.UserCredential)(int64,bool,error){
	userId,valid,err:=a.db.ValidateCredential(ctx,input)
	if err!=nil{
		return 0,false,err
	}
	return userId,valid,nil
}