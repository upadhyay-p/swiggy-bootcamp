package main

import
(
	"fmt"
	"time"
)

func main(){
	t := time.Date(2020, time.July, 22, 8, 5,7,5,time.UTC)
	fmt.Printf("%v/%v/%v\n", int(t.Month()), t.Day(), t.Year())

	now:=time.Now()
	fmt.Printf("Time now %s\n",now)

	fmt.Println("Day ",t.Day())

}