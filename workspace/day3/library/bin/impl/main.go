package main

import (
	"../stringutil"
	"fmt"
)

func main(){
	full, length := stringutil.FullNameNakedReturn("Priya","U")
	fmt.Println(full, length)
}