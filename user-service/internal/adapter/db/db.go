package db

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

type Adapter struct {
	db *sql.DB
}
var connectionString = os.Getenv("CONN_STR")
func NewAdapter(db *sql.DB)(*Adapter,error){
	db,err:=sql.Open("postgres",connectionString)
    if err!=nil{
		return &Adapter{},err
	 }
	 return &Adapter{db: db},nil
 }

