package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is string")
	case int:
		fmt.Println("Value", value, "is Int")
	default:
		fmt.Println("Unknown type")
	}
}
