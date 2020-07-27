package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"google.golang.org/grpc"
	"net/http"
	"gin-gRPC/greet/pb"
)

func GetCustomers(c *gin.Context) {
	req := &pb.CustomerRequest{
		Cust:&pb.Customer {
			CustomerID:"101",
			CustomerName:"Priya",
			OrderReview:"Good",
		},
	}

	res, err := customerServiceClient.CustomerInfo(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res.Result),
	})
}

var customerServiceClient pb.CustomerServiceClient

func main(){
	fmt.Println("Hi, I'm in client..")
	conn ,err := grpc.Dial("localhost:5051",grpc.WithInsecure())

	if err!=nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
	}

	defer conn.Close()

	customerServiceClient = pb.NewCustomerServiceClient(conn)

	// setting up gin
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/customers",GetCustomers)

	if err := router.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
