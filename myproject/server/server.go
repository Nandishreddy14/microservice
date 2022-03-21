package main

import (
	"context"
	"fmt"
	pb "golang/microservices/myproject/gen/proto"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pub "golang/microservices/myproject/publisher"

	sub "golang/microservices/myproject/subscriber"
)

type TestapiServer struct {
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (server *TestapiServer) Getdetails(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	  pub.Publish()

	  go sub.Subscribe()

	return &pb.UserResponse{Name: "Nandish Narayanareddy"}, nil

}
func main() {

	go func() {

		mux := runtime.NewServeMux()

		pb.RegisterTestapiHandlerServer(context.Background(), mux, &TestapiServer{})

		err := http.ListenAndServe(":9091", mux)

		if err != nil {

			fmt.Println("http servre is not running", err)
		}

	}()

	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Println(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestapiServer(grpcServer, &TestapiServer{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}
