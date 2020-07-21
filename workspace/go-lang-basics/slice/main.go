package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = [] string {"red", "blue", "green"}

	fmt.Println(colors)
	colors = append(colors, "Yellow")
	fmt.Println(colors)
	colors = colors[:len(colors)-1]

	fmt.Println(colors)

	ints := make([]int, 5, 5)
	fmt.Println("Initial capacity: ", cap(ints))
	ints[1]=11
	ints[0]=99
	ints[2]=22
	ints[3]=44
	ints[4]=55
	fmt.Println("After adding 5 elems capacity: ", cap(ints))
	fmt.Println(ints)
	ints = append(ints,88)
	fmt.Println(ints)
	fmt.Println("Capacity now ",cap(ints))

	sort.Ints(ints)
	fmt.Println("Ascending order--", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
}