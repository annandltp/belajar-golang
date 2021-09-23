package main

import "fmt"

type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
