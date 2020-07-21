package main

import "fmt"

func main() {
	var p *int

	var v int = 44
	p = &v

	if p != nil {
		fmt.Println("Value of p: ", *p)
	} else {
		fmt.Println("P is nil")
	}

	var val1 float64 = 33.2
	pointerF := &val1

	fmt.Println("Val1 : ",*pointerF)

}
