package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func GetFirstName(person Person) string {
	return person.FirstName
}
func GetLastName(person Person) string {
	return person.LastName
}

func GetAge(person Person) int {
	return person.Age
}

func setFirstName(person Person, firstName string) {
	person.FirstName = firstName
}

func setLastName(person Person, lastName string) {
	person.LastName = lastName
}

func setAge(person Person, age int) {
	person.Age = age
}
