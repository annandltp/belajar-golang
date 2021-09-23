package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Anan"
	middleName = "dela"
	lastName = "titis"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
