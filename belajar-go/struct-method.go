package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHo() {
	fmt.Println("Ho Fron", a.Address)
}

func main() {
	var eko Customer
	eko.Name = "Anan"
	eko.Address = "Bekasi"
	eko.Age = 50

	eko.sayHi("Anan")
	eko.sayHo()
}
