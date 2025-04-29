package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	
	"time"

	"github.com/kida21/userservice/internal/application/core/domain"
	
)

var (
	ErrDuplicate = errors.New("duplicate Email")
)
type Adapter struct {
	db *sql.DB
}

func NewAdapter(db *sql.DB)(*Adapter){
  return &Adapter{db: db}
 }

func (a *Adapter)Insert(ctx context.Context, user *domain.UserModel)(bool,error){
	query:=`
	        INSERT INTO users(firstname,lastname,email,password_hash) 
	        VALUES($1,$2,$3,$4) RETURNING id,version,created_at
			`
	args:=[]any{user.FirstName,user.LastName,user.Email,user.Password.Hash}
	ctx,cancel:=context.WithTimeout(ctx ,time.Second * 9)
	defer cancel()
	//for debugging purpose
	defer log.Print("user:",user.FirstName+user.LastName)
	err:=a.db.QueryRowContext(ctx,query,args...).Scan(&user.Id,&user.Version,&user.Creadted_at)
	if err!=nil{
         switch{
		 case err.Error() == `pq:duplicate key value violates unique constraint user_email key`:
			return false,ErrDuplicate
		 default:
			return false,err
		 }
	}
	return true,nil
 }
