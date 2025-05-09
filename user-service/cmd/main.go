package main

import (
	"log"
	"net"

	"github.com/kida21/userservice/internal/adapter/db"
	"github.com/kida21/userservice/internal/adapter/grpc"
	"github.com/kida21/userservice/internal/application/core/api"
	Grpc"google.golang.org/grpc"
	//pb "github.com/kida21/userservice/gen"
	pb"github.com/kida21/Microservice-With-Hexagonal-architecture/z-proto/user"
)

func main() {
	conn,err:=db.OpenConnection()
	if err!=nil{
		log.Fatal(err)
	}
	adapter:=db.NewAdapter(conn)
	api:=api.NewApplication(adapter)
	handler:=grpc.NewHandler(api)

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