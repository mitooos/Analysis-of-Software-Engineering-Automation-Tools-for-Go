package main

import (
	"context"
	"fmt"
	pb "grpc-example/protos"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTimeZoneServer
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTimeZoneServer(grpcServer, &server{})

	log.Printf("listening on address: %q", address)
	grpcServer.Serve(lis)
}

func (*server) WhatTimeIsIt(ctx context.Context, req *pb.TimeRequest) (*pb.TimeReply, error) {
	log.Printf("request: %+v", req)
	loc, err := time.LoadLocation(req.TimeZone)
	if err != nil {
		return &pb.TimeReply{}, fmt.Errorf("please provide a valid timezone")
	}
	now := time.Now().In(loc).Format("2006-01-02 15:04:05")
	return &pb.TimeReply{Time: now}, nil
}
