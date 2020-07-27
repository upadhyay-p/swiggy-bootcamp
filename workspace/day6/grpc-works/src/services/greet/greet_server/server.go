package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"services/greet/greetpb"
)

type Server struct {

}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Println("Function called... ")

	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()

	result := "Hello : "+fname+ ", "+lname

	res := &greetpb.GreetResponse{
		Result:result,
	}
	return res,nil

}

func (*Server) GreetFullName(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetFullNameResponse, error) {
	fmt.Println("GreetFullName called... ")
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()

	result := strings.ToUpper(fname+ " "+lname)

	res := &greetpb.GreetFullNameResponse{
		Greeting:result,
	}
	return res,nil
}

func main() {
	fmt.Println("Hello from server.")

	lis, err := net.Listen("tcp","0.0.0.0:5051")
	if err!=nil {
		log.Fatalf("Sorry failed to load server %v:",err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &Server{})

	if s.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve %v",err)
	}

}