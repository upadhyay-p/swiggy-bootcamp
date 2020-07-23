package main

import "fmt"

func main() {
	DoSomeWork()
	fmt.Println(AddValues(2,3))
	fmt.Println(AddAll(2,3,5))
	f,l:=FullName("Priya","Upadhyay")
	fmt.Printf("Full name is %s, length is %v\n",f,l)
	full,length := FullNameNakedReturn("Harry","Styles")
	fmt.Printf("Full name is %s, len is %v",full,length)
}

func DoSomeWork() {
	fmt.Println("Doing your work")
}

func AddValues(v1 int, v2 int) int {
	return v1+v2
}
//go does not have function overloading

func AddAll(values ... int) int {
	sum:=0
	for i:=range values {
		sum+=values[i]
	}
	return sum
}
//return fullname and length
func FullName(f,l string) (string,int) {
	full:=f+" "+l
	length:= len(full)
	return full,length
}

//you can name the return variable
//in theses kinds of functions we are not returning
//these are called Naked Functions since they don't return anything
func FullNameNakedReturn(f,l string)(full string, length int){
	//full, length are already declared, hence no : is required
	full = f+" "+l
	length = len(full)
	return
}