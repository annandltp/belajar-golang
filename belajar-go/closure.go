package main

import "fmt"

func main() {
	name := "Anan"
	counter := 0

	increment := func() {
		name = "dela"
		fmt.Println("Increment")
		counter++
	}

	increment()
	// increment()

	fmt.Println(counter)
	fmt.Println(name)
}
