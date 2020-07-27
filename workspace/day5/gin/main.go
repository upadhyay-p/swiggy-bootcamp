package main


// we need who can generate RESTful data
// we have installed gin for it

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Order struct {
	OrderID int `json: "orderId"`
	CustomerName string `json:"custName"`
	OrderReview string `json:"review"`
}


func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello GIN",
	})
}

var orders = []Order{
	{101,"C1","Good food"},
	{102,"C2","Good food"},
	{103,"C3","Loved the food"},

}

func GetOrders(c *gin.Context){
	c.JSON(200, &orders)
}

func GetSpecificOrder(c *gin.Context){
	orderId := c.Param("orderId")
	order := c.Param("order")

	c.JSON(200, gin.H{
		"orderId":orderId,
		"order":order,
	})
}

func PostOrder(c *gin.Context){
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	if err!=nil {
		fmt.Println("Sorry no content found: ",err.Error())
	}
	fmt.Println(string(content))

	c.JSON(http.StatusCreated, gin.H{
		"message":string(content),
	})

	//we have to append it to slice DIY
}
//localhost:9000/api/orders?minNoOfOrders=100&showOrder=asc
func GetSpecificOrdersByQuery(c *gin.Context){
	minNumberOfOrder := c.Query("minNoOfOrders")
	showOrder := c.Query("showOrder")

	c.JSON(200, gin.H{
		"minNoOfOrders":minNumberOfOrder,
		"showOrder":showOrder,
	})
}

func main(){
	// router:=gin.New() -->this will not carry middle tier work

	router := gin.Default()

	api := router.Group("/api")

	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})


	//http://localhost:9000/api/
	api.GET("/",HomePage)

	//http://localhost:9000/api/orders
	api.GET("/orders",GetOrders) //to return all orders

	api.POST("/order",PostOrder)

	//colon indicates path parameter
	//http://localhost:9000/api/orders/101/asc
	api.GET("/orders/:orderId/:order", GetSpecificOrder) //to return 1 order

	//query params
	//localhost:9000/api/orders?minNoOfOrders=100&showOrder=asc
	api.GET("/query/orders/",GetSpecificOrdersByQuery)

	router.Run("localhost:9000")
}