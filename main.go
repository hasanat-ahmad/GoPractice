package main

import "fmt"

// import "fmt"

func main() {
	var arr [5]int

	InsertElement(arr[:], 5)
	InsertElement(arr[:], 7)
	InsertElement(arr[:], 8)
	InsertElement(arr[:], 9)
	InsertElement(arr[:], 9)

	UpdateElement(arr[:], 8, 6)
	fmt.Println(arr)
	

}
