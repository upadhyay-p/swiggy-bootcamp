package main

import (
	"fmt"
)

func main() {
	//Defer---this will be executed last in its scope
	//Defer works in LIFO order
	defer fmt.Println("Connection Close")
	fmt.Println("Connection Open")
	myFunc()
	defer fmt.Println("Testing defer")
	defer fmt.Println("Testing defer1")

	fmt.Println("Do the work")
	//Some problem here
	panic("SomeError")

}


func myFunc() {
	defer fmt.Println("I'm in defer MyFunc")
	fmt.Println("Not defer in MyFunc")
}