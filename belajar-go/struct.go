package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Anan"
	eko.Address = "Bekasi"
	eko.Age = 50

	fmt.Println(eko)

	// var joko Customer = Customer
	joko := Customer{
		Name:    "Anan",
		Address: "Bekasi",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Anan", "Jakarta", 12}
	fmt.Println(budi)
}
