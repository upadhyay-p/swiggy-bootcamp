package main

import "fmt"

func main() {
	dict := make(map[string] string)
	fmt.Println(dict)

	dict["B"]="1"
	dict["A"]="2"
	dict["C"]="3"

	for key,val := range dict{
		fmt.Printf("%v --> %v\n",key,val)
	}

	delete(dict,"A")
	fmt.Println(dict)
}