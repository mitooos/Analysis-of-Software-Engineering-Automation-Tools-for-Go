package main

import (
	"context"
	pb "grpc-example/protos"
	"log"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	address := "0.0.0.0:50051"
	cc, err := grpc.Dial(address, opts)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	client := pb.NewTimeZoneClient(cc)

	resp, err := client.WhatTimeIsIt(context.Background(), &pb.TimeRequest{TimeZone: "America/New_York"})
	if err != nil {
		panic(err)
	}

	log.Printf("response: %+v", resp)
}
