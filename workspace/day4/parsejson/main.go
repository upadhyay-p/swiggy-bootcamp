package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Posts struct {
	Title string
	UserID int
}

func main () {
	url := "https://jsonplaceholder.typicode.com/posts"
	content := ContentFromServer(url)
	fmt.Println(PostFromJson(content))
}

func PostFromJson(content string) [] Posts {
	posts := make([]Posts, 0 ,20)
	decoder := json.NewDecoder(strings.NewReader(content))

	_,err := decoder.Token()
	CheckError(err)

	var post Posts

	for decoder.More() {
		err := decoder.Decode(&post)
		CheckError(err)
		posts = append(posts, post)
	}

	return posts
}

func ContentFromServer(url string) string {
	res,err := http.Get(url)
	CheckError(err)
	defer res.Body.Close()
	bytes,err := ioutil.ReadAll(res.Body)
	CheckError(err)
	return string(bytes)
}

func CheckError(err error) {
	if err!=nil {
		panic(err)
	}
}