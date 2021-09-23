package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Eko"
	names[1] = "Anan"
	names[2] = "Yoyo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		89,
		87,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}
