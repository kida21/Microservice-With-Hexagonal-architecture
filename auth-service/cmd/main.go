package main

import (
	"log"
	"net"
	"os"

	handler "github.com/kida21/authservice/grpc"
	//pb "github.com/kida21/authservice/grpc/gen"(replaced with published module of github.com/kida21)
	"github.com/kida21/authservice/internal/adapter"
	"github.com/kida21/authservice/internal/application/core/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authpb"github.com/kida21/Microservice-With-Hexagonal-architecture/z-proto/auth"
)

func main() {

    connString:=os.Getenv("USERSVR_ADDR")
	                                          //insecure for development purpose only
	conn,err:=grpc.NewClient(connString,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("unable to make a client connection :%v",err)
	}
    adapter:=adapter.NewAdapter(conn)
	service:=api.NewService(adapter)
	handler:=handler.NewHandler(service)

    listener,err:=net.Listen("tcp",":50051")
	if err!=nil{
       log.Fatal(err)
	}
	log.Println("Server running on :50051")
    grpcServer:=grpc.NewServer()
	authpb.RegisterAuthenticationServer(grpcServer,handler)
 	if err=grpcServer.Serve(listener);err!=nil{
		log.Fatal(err)
	}
	
	


}