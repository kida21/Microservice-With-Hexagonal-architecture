package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/kida21/userservice/internal/application/core/domain"
	_ "github.com/lib/pq"
)

var (
	ErrDuplicate = errors.New("duplicate Email")
)
type Adapter struct {
	db *sql.DB
}

func NewAdapter(db *sql.DB)(*Adapter,error){
  return &Adapter{db: db},nil
 }

 func (a Adapter)Insert(ctx context.Context,user *domain.UserModel)(bool,error){
	query:=`INSERT INTO users(firstname,lastname,email,password) 
	         VALUES($1,$2,$3,$4) RETURNING id,version,created_at`
	ctx,cancel:=context.WithTimeout(ctx ,time.Second * 5)
	args:=[]any{user.FirstName,user.LastName,user.Email,user.Password}
	defer cancel()
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
