package main

import (
	"flag"
	"fmt"
)

func main() {
	// var host *string = flag.String("host", "localhost", "Put your host")
	// var user *string = flag.String("user", "root", "Put your database user")
	// var password *string = flag.String("host", "localhost", "Put your database password")
	// var number *int = flag.Int("number", 100, "Put your number")

	// flag.Parse()

	// fmt.Println("Host: *", *host)
	// fmt.Println("User: *", *user)
	// fmt.Println("Password: *", *password)
	// fmt.Println("Number: *", *number)

	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your user database")
	pass := flag.String("pass", "", "Put your password database")

	flag.Parse()

	fmt.Println(*host, *user, *pass)
	fmt.Println("Host: *", *host)
	fmt.Println("User: *", *user)
	fmt.Println("Password: *", *pass)
}
