package main

import "fmt"

func main() {
	name := "Anan123"

	switch name {
	case "Anan":
		fmt.Println("Hallo Anan")
	case "Joko":
		fmt.Println("Hallo Joko")
	default:
		fmt.Println("Hallo")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
