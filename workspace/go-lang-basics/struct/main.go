package main

import "fmt"

type Item struct{
	Name string
	Price int

}

func main() {
	item1 := Item{Name:"Soap", Price:10}
	fmt.Println(item1)
}