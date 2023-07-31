package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sudhagayathri/golang/grpc"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("couldnt greet ", err)
	}
	log.Printf(res.Message)

}
