package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	res,err := http.Get(url)
	if err!=nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Printf("Content type: %T",res)
	//fmt.Println(res.Body)

	bytes,err := ioutil.ReadAll(res.Body)
	if err!=nil {
		panic(err)
	}
	content := string(bytes)
	fmt.Println(content)
	fmt.Println(res.StatusCode)

}
