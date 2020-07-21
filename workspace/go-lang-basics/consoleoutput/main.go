package main

import "fmt"

func main() {
	str1 := "String 1"
	str2 := "String 2"
	str3 := "Strring 3"
	aNumber := 33
	importsTrue := true

	strLength, _ := fmt.Println(str1,str2,str3)
	fmt.Println("The string length is ",strLength)
	fmt.Printf("Value of a number %v\n", aNumber)
	fmt.Printf("Value of boolean %v\n", importsTrue)
	fmt.Printf("Value as float %.2f", float64(aNumber))
	fmt.Printf("\nData types %T ,%T, %T", str1, aNumber, importsTrue)
	myString := fmt.Sprintf("\nData types %T ,%T, %T", str1, aNumber, importsTrue)
	fmt.Println(myString)
}