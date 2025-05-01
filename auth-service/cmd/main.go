package main

import (
	"log"
	"net"
	"os"

	handler "github.com/kida21/authservice/grpc"
	pb "github.com/kida21/authservice/grpc/gen"
	"github.com/kida21/authservice/internal/adapter"
	"github.com/kida21/authservice/internal/application/core/api"
	"google.golang.org/grpc"

	
)

func main() {
	adapter:=adapter.NewAdapter()
	service:=api.NewService(adapter)
	handler:=handler.NewHandeler(service)

    listener,err:=net.Listen("tcp",":50051")
	if err!=nil{
       log.Fatal(err)
	}
	log.Println("Server running on :50051")
    grpcServer:=grpc.NewServer()
	pb.RegisterAuthenticationServer(grpcServer,handler)
 	if err=grpcServer.Serve(listener);err!=nil{
		log.Fatal(err)
	}
	connString:=os.Getenv("USERSVR_ADDR")
	_,err=grpc.NewClient(connString)
	if err!=nil{
		log.Fatalf("unable to make a client connection :%v",err)
	}
	

}