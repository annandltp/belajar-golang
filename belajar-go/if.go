package main

import "fmt"

func main() {
	var name = "Joko1111"

	if name == "Anan" {
		fmt.Println("Hallo")
	} else if name == "Joko" {
		fmt.Println("Hi, Joko")
	} else {
		fmt.Println("Hi, ")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nmaa sudah benar")
	}
}
