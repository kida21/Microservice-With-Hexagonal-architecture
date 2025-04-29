package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)
type Password struct{
	plaintext *string
	Hash []byte
}
type UserModel struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    Password `json:"-"`
	Version     int64  `json:"-"`
	Activated   bool   `json:"activated"`
	Creadted_at time.Time `json:"created_at"`
}
func (p *Password) Set(plaintext string)error{
	hash,err:=bcrypt.GenerateFromPassword([]byte(plaintext),bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	p.plaintext=&plaintext
	p.Hash=hash

	return nil
}
func (p *Password) Matches(plaintext string)(bool,error){
	err:=bcrypt.CompareHashAndPassword(p.Hash,[]byte(plaintext))
	if err!=nil{
		switch{
		   case errors.Is(err,bcrypt.ErrMismatchedHashAndPassword):
			return false,err
		   default:
			return false,err
		}
	}
	return true,nil
}