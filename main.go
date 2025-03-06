package main

func main() {
	var arr [5]int

	InsertElement(arr[:], 10)
	InsertElement(arr[:], 5)
	InsertElement(arr[:], 71)
	InsertElement(arr[:], 8)
	InsertElement(arr[:], 9)

	FindingMaxInArray(arr[:])
	FindingMinInArray(arr[:])
	// fmt.Println(arr)

}
