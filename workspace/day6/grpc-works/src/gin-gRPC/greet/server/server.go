package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gin-gRPC/greet/pb"
)

type server struct{}

func (*server) CustomerInfo(ctx context.Context, req *pb.CustomerRequest) (*pb.CustomerResponse, error){
	fmt.Println("Function called... ")

	custID := req.GetCust().GetCustomerID()
	custName := req.GetCust().GetCustomerName()
	custRev := req.GetCust().GetOrderReview()

	result := custID+ ", "+custName +", "+custRev

	res := &pb.CustomerResponse{
		Result:result,
	}
	return res,nil

}

func main() {
	fmt.Println("Hello from grpc server.")

	lis, err := net.Listen("tcp","0.0.0.0:5051")
	if err!=nil {
		log.Fatalf("Sorry failed to load server %v:",err)
	}

	s := grpc.NewServer()

	pb.RegisterCustomerServiceServer(s, &server{})


	if s.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve %v",err)
	}

}
