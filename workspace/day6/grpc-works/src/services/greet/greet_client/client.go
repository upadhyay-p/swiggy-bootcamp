package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"services/greet/greetpb"
)

func callGreet(c greetpb.GreetServiceClient) {
	fmt.Println("In Call greet fucntion...")

	req := &greetpb.GreetRequest{
		Greeting:&greetpb.Greeting {
			FirstName:"Priya",
			LastName:"Upadhyay",
		},
	}

	res,err := c.Greet(context.Background(), req)

	if err!=nil {
		log.Fatalf("Error while calling greet: %v", err)
	}

	log.Println("Response from the greet is - ",res.Result)


	res1 ,err := c.GreetFullName(context.Background(), req)

	if err!=nil {
		log.Fatalf("Error while calling greet: %v", err)
	}

	log.Println("Response from the greetFullName is - ",res1.Greeting)

}

func main() {
	fmt.Println("Hi, I'm in client..")
	conn ,err := grpc.Dial("localhost:5051",grpc.WithInsecure())

	if err!=nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	callGreet(c)
}
