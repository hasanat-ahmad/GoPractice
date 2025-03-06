package main

import "fmt"

func main() {
	var password string = "Hasahsahsahs"
	LengthCheck(password)
	ContainsUpperCase(password)
	ConatainsLowerCase(password)
	ContainSpecialChar(password)

	if LengthCheck(password) && ContainsUpperCase(password) && ConatainsLowerCase(password) && ContainSpecialChar(password) {
		fmt.Print("Strong password")
	} else {
		fmt.Println("Weak password")
	}
}
