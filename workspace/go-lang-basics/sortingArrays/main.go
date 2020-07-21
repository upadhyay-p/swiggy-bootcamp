package main

import "fmt"

func main() {
	var name [5] string
	name[0] = "Priya"
	name[1] = "Bhavi"
	name[2] = "Shashi"
	name[3] = "Varun"
	name[4] = "Jyotish"
	fmt.Println(name)

	var numbers = [3] int {1,2,3}

	fmt.Println((len(numbers)))
}