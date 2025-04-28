package main

import (
	"log"
	"net"

	"github.com/kida21/userservice/internal/adapter/db"
	"github.com/kida21/userservice/internal/adapter/grpc"
	"github.com/kida21/userservice/internal/application/core/api"
	Grpc"google.golang.org/grpc"
	pb "github.com/kida21/userservice/internal/adapter/grpc/gen"
)

func main() {
	conn,err:=db.OpenConnection()
	if err!=nil{
		log.Fatal(err)
	}
	repo:=db.NewAdapter(conn)
	service:=api.NewApplication(repo)
	handler:=grpc.NewHandler(service)

	listner,err:=net.Listen("tcp",":50021")
	if err!=nil{
		log.Fatal(err)

	}
	log.Println("server started on:50021")
	grpcserver:=Grpc.NewServer()
    pb.RegisterUserServer(grpcserver,handler)
	if err=grpcserver.Serve(listner);err!=nil{
		log.Println(err)
		log.Fatal(err)
		
	}
	

}