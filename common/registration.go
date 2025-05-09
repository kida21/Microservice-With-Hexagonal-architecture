package common

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)
type Serviceregistrar struct{
   Client *api.Client
}

func NewServiceRegistrar(addr string)(*Serviceregistrar){
   config:= api.DefaultConfig()
   config.Address=addr
   client,err:=api.NewClient(config)
   if err!=nil{
	  log.Fatal(err)
   }
   return &Serviceregistrar{Client: client}
}
func generateServiceId(servicName string)(string){
  id:=uuid.New().String()
  return fmt.Sprintf("%s-%s",servicName,id)
}

func (s*Serviceregistrar) Register(serviceName,address string,port int)(string,error){
   serviceId:=generateServiceId(serviceName)
   registration:=&api.AgentServiceRegistration{
	ID: serviceId,
	Name: serviceName,
	Port: port,
	Address: address,
    Check: &api.AgentServiceCheck{
		GRPC: fmt.Sprintf("%s:%d",address,port),
		Interval: "10s",
		Timeout: "3s",
		DeregisterCriticalServiceAfter: "2m",
	},
   } 
   log.Printf("Registering a service on the address:%sport:%d",address,port)
   err:=s.Client.Agent().ServiceRegister(registration)
   if err!=nil{
	log.Fatalf("failed to register a service:%v",err)
   }
   return serviceId,nil
}

func (s *Serviceregistrar) DeregisterService(serviceId string)(error){
   if err:=s.Client.Agent().ServiceDeregister(serviceId);err!=nil{
	log.Fatalf("failed to deregister a service:%v",err)
   }
   return nil
}