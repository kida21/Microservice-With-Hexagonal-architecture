package api

import (
	"github.com/kida21/userservice/internal/application/core/domain"
	"github.com/kida21/userservice/internal/application/ports"
)


type Application struct {
	db ports.DBPort
}
func NewApplication(db ports.DBPort)(*Application){
   return &Application{db: db}
}

func (a *Application) Register(user domain.UserModel)( domain.UserModel,error){
	user,err:=a.db.Insert(&user)
	if err!=nil{
		return domain.UserModel{},err
	}
	return user,nil
}