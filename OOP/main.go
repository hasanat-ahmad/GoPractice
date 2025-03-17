package main

import "fmt"

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	fmt.Println(GetFirstName(person))
}