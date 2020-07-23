package main

import "fmt"

var isConnected bool =false
func main() {
	fmt.Printf("Connection open: %v\n",isConnected)
	businessLogic()
	fmt.Printf("Connection Open: %v\n",isConnected)

}

func businessLogic() {
	defer disconnect()
	connect()
	fmt.Println("DB connection status businessLogic :",isConnected)
	fmt.Println("Business logic goes here")
}

func connect(){
	isConnected=true
	fmt.Println("DB connection is Open")

}

func disconnect() {
	isConnected=false
	fmt.Println("DB connection is close")
}