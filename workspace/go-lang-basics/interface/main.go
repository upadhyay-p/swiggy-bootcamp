package main

import "fmt"

type Animal interface{
	Speak() string
}

func (e Emp) Speak() string {
	return "Hello... Hello..."
}
func (d Dog) Speak() string {
	return "Wof.. Wof... "
}

func (c Cat) Speak() string {
	return "Meow.. Meow... "
}

// structure or noun
type Dog struct {}
type Emp struct {}
type Cat struct {}



func main() {
	dog1 := Animal(Dog{})
	fmt.Println(dog1)
	fmt.Println(dog1.Speak())


	animals := [] Animal{Emp{}, Dog{}, Cat{}}

	for _, animal := range  animals {
		fmt.Println(animal.Speak())
	}


}