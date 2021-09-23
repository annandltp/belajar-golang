package main

import "fmt"

func getFullName() (string, string, string) {
	return "Anan", "Dela", "Yoyo"
}

func main() {
	// firstName, middleName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
