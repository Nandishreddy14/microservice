package main

import (
	"context"
	"fmt"

	pb "golang/microservices/myproject/gen/proto"

	"google.golang.org/grpc"
)



func main() {

	conn,err := grpc.Dial("localhost:9090",grpc.WithInsecure())

	if err!=nil{
		fmt.Println("error connecting to grpc server")
	}
    
	fmt.Println("testing")

	client:=pb.NewTestapiClient(conn)

	resp,err:=client.Getdetails(context.Background(),&pb.UserRequest{Id:509772})
   
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(resp)
	 
}