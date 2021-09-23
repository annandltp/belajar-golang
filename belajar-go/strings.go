package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Anandela", "Anan"))
	fmt.Println(strings.Contains("Anandela", "Andi"))

	fmt.Println(strings.Split("Anandela Titis", " "))

	fmt.Println(strings.ToLower("Anandela Titis"))
	fmt.Println(strings.ToUpper("Anandela Titis"))
	fmt.Println(strings.ToTitle("Anandela Titis"))
	fmt.Println(strings.Trim("   Anandela Titis    ", " "))
	fmt.Println(strings.ReplaceAll("Anan Dian", "Anan", "Toti"))
}
