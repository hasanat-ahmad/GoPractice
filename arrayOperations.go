package main

// import "fmt"

func InsertionInArray(arr[] int, number int){
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0{
			arr[i] = number
		}
	}
}

func DeleteElementInArray(arr[] int, number int){
	for i := 0; i < len(arr); i++ {
		if (arr[i] == number){
			arr[i] = 0
		}
	}
}