package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface {
	Eat() string
	Sleep() string
	Breath() string
}

type Addr struct {
	Hno    string
	Street string
	City   string
	Pin    string
}

type Emp struct {
	EmpId int64
	Name  string
	Sal   int64
	Eaddr Addr
}

func (Emp) Eat() string {
	return "Eats"
}

func (Emp) Sleep() string {
	return "Sleeps"
}

func (Emp) Breath() string {
	return "Breaths"
}

func Print(e Emp) {
	fmt.Printf("EmpID: %v, Name: %s, Sal: %v, Address: %s, %s, %s, %s\n\n", e.EmpId, e.Name, e.Sal, e.Eaddr.Hno, e.Eaddr.Street, e.Eaddr.City, e.Eaddr.Pin)

}

func main() {
	fmt.Println("Hello, Welcome!")

	emps := make([]Emp, 0, 5)

	for true {
		fmt.Println("Please choose a number 1 or 2")
		fmt.Println("1. Enter Employee Details")
		fmt.Println("2. Quit")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		if "2" == strings.TrimSpace(choice) {
			//for elem := range emps{
			//	Print(elem)
			//}
			fmt.Println(emps)
			break
		}
		var e Emp
		fmt.Print("Employee ID: ")
		eid, _ := reader.ReadString('\n')
		e.EmpId, _ = strconv.ParseInt(strings.TrimSpace(eid), 10, 64)

		fmt.Print("\nName: ")
		name, _ := reader.ReadString('\n')
		e.Name = strings.TrimSpace(name)

		fmt.Print("\nSalary: ")
		sal, _ := reader.ReadString('\n')
		e.Sal, _ = strconv.ParseInt(strings.TrimSpace(sal), 10, 64)

		var addr Addr
		fmt.Print("\nAddress\nHouse No.: ")
		hno, _ := reader.ReadString('\n')
		addr.Hno = strings.TrimSpace(hno)

		fmt.Print("\nStreet: ")
		street, _ := reader.ReadString('\n')
		addr.Street = strings.TrimSpace(street)

		fmt.Print("\nCity: ")
		city, _ := reader.ReadString('\n')
		addr.City = strings.TrimSpace(city)

		fmt.Print("\nPin: ")
		pin, _ := reader.ReadString('\n')
		addr.Pin = strings.TrimSpace(pin)

		e.Eaddr = addr

		emps = append(emps, e)
		fmt.Printf("EmpID: %v, Name: %s, Sal: %v, Address: %s, %s, %s, %s\n\n", e.EmpId, e.Name, e.Sal, e.Eaddr.Hno, e.Eaddr.Street, e.Eaddr.City, e.Eaddr.Pin)
		fmt.Println(e.Eat())
		fmt.Println(e.Sleep())
		fmt.Println(e.Breath())
		fmt.Println()
	}

}
