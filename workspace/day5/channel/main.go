package main

import (
	"fmt"
	"strconv"
)

func main() {


	c := make(chan string)
	c1 := make(chan string)
	go DoMyWork("A", c)
	go DoMyWork1("B", c1)
	for !IsClosed(c) {
		fmt.Println(<-c)
	}
	//for msg := range c {
	//	//msg = <-c
	//	fmt.Println(msg)
	//}
	close(c)
}

func DoMyWork(name string, c chan string) {
	for i :=1; i<=5; i++ {
		c <- "Doing work of "+ name+ " executed "+strconv.Itoa(i)
		//time.Sleep(time.Microsecond*500)
	}
	//close(c)
}

func DoMyWork1(name string, c chan string) {
	for i :=1; i<=5; i++ {
		c <- "Doing work of "+ name+ " executed "+strconv.Itoa(i)
		//time.Sleep(time.Microsecond*500)
	}
	//close(c)
}

func IsClosed (ch chan string) bool {
	select {
		case<-ch :
			return true
		default:
	}
	return false
}

