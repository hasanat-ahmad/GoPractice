package main

import "fmt"

func FindingMaxInArray(arr []int) int {
	
	var max int = arr[0]
	for i := 0; i < len(arr); i++{
		if arr[i] > max{
			max = arr[i]
		}
	}
	fmt.Println("Max in array is", max)
	return max
}
