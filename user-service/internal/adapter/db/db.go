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
	ErrNotFound = errors.New("sql no rows")
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

  func (a *Adapter)Update(ctx context.Context,user *domain.UserModel)(int64,int64,error){
	query:=`UPDATE users 
	        SET firstname = $1,lastname = $2,email= $3,
			password_hash = $4,version = version + 1 WHERE id = $5 AND version = $6 RETURNING id,version`
	 
	  args:=[]any{user.FirstName,user.LastName,user.Email,user.Password.Hash,user.Id,user.Version}
	  ctx,cancel:=context.WithTimeout(ctx,time.Second * 9)
	  defer cancel()
	  defer log.Println("userid:",user.Id,"version:",user.Version)
	 err:=a.db.QueryRowContext(ctx,query,args...).Scan(&user.Id,&user.Version)
	 if err!=nil{
		return 0,0,err
	 }
	 return user.Id,user.Version,nil
  }

 func(a * Adapter)ValidateCredential(ctx context.Context,input *domain.UserCredential)(int64,bool,error){
     var user domain.UserModel
	query:=`SELECT id,password_hash,email FROM users WHERE email = $1`
	ctx,cancel:=context.WithTimeout(ctx,time.Second*9)
	defer cancel()
	err:=a.db.QueryRowContext(ctx,query,input.Email).Scan(&user.Id,&user.Password.Hash,&user.Email)
	if err!=nil{
		switch{
		case errors.Is(err,sql.ErrNoRows):
			return 0,false,err
		default:
			return 0,false,err
		}
	}
	 valid,err:=input.Compare(input.Password,user.Password.Hash)
	 if err!=nil{
		return 0,false,err
	 }
	 log.Println(user.Id,":",valid)
	 return user.Id,valid,nil
}
