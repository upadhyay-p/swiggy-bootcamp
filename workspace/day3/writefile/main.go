package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from the other side"

	file,err := os.Create("./day3/writefile/sample.txt")
	checkError(err)

	defer file.Close()
	l, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Saved %v chars to file", l)

	bytes := []byte(content)
	err = ioutil.WriteFile("./day3/writefile/inbytes.txt", bytes, 0644)
	checkError(err)
}

func checkError(err error) {
	if err!=nil {
		panic(err)
	}
}