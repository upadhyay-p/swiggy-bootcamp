package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name-")
	name,_ := reader.ReadString('\n')
	fmt.Println("Your name is ",name)
	fmt.Println("Enter salary: ")
	sal,_:=reader.ReadString('\n')
	fmt.Println("Enter months: ")
	month,_:=reader.ReadString('\n')
	fSal,_ := strconv.ParseFloat(strings.TrimSpace(sal),64)
	fMon,_ := strconv.ParseFloat(strings.TrimSpace(month), 64)
	fmt.Printf("Total: %.2f", fSal*fMon)

}