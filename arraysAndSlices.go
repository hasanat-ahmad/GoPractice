package main

import "fmt"

func Array() {
	var marks [5]int = [5]int{1, 2, 4, 5}
	fmt.Println(marks)
	fmt.Println("Length of array is", len(marks))
	fmt.Println("Capacity of array is", cap(marks))
}

func slices() {
	var marks = []int{1, 2, 4, 5, 0}
	fmt.Println("Length of array is", len(marks))
	fmt.Println("Capacity of array is", cap(marks))
	fmt.Println(marks)
}
