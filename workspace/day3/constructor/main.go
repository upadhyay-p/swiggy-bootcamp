package main

import (
	"fmt"
)

type Trade struct {
	CompanyName string
	Value int
}

func NewTrade(Name string, Val int) (*Trade, error) {
	if Name == "" {
		return nil, fmt.Errorf("Company name can't be empty")
	} else {
		return &Trade{CompanyName:Name, Value:Val},nil
	}
}

func main() {
	t1,err := NewTrade("",0)
	if err==nil {
		fmt.Println(*t1)
	} else {
		fmt.Println(err)
	}
	t2,err := NewTrade("Swig",122)
	if err==nil {
		fmt.Println(*t2)
	} else {
		fmt.Println(err)
	}

}

