package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := [] string {
		"https://golang.org", //string
		"https://jsonplaceholder.typicode.com/posts", //json
		"https://httpbin.org/xml", //xml
	}

	var wg sync.WaitGroup
	for _,url := range urls {
		wg.Add(1)
		//returnType(url)
		//anonymous function
		go func(url string) {
			returnType(url)
			wg.Done()
		} (url)

		//alternate fashion of writing above syntax/goroutine
		//go returnType(url)

	}
	wg.Wait()
}


func returnType(url string) {
	res,err := http.Get(url)
	CheckError(err)
	defer res.Body.Close()
	ctype := res.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)

}

func CheckError(err error) {
	if err!=nil {
		panic(err)
	}
}
