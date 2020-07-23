package main

import (
	"fmt"
	"net/http"
)

//type Handler interface {
	//ServeHTTP(ResponseWriter, *Request)
//}

type Hello struct {

}

func (h Hello) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "<html><h2>Hey</h2></html>")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:9000", h)
	CheckError(err)
}

func CheckError(err error) {
	if err!=nil {
		panic(err)
	}
}
