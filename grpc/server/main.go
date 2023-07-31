package main

import (
	"log"
	"net"

	pb "grpc"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start ", err)
	}
	grpcserver := grpc.NewServer() //creates grpc instance
	pb.RegisterGreetServiceServer(grpcserver, &helloServer{})
	//Above, Registers the server implementation (helloServer) with the gRPC server. It associates the methods defined in the helloServer struct with the corresponding gRPC service methods.
	log.Println("server started at ", lis.Addr())
	if err := grpcserver.Serve(lis); err != nil { // Starts the gRPC server
		log.Fatalf("failed to start ", err)
	}
}
