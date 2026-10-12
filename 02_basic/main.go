package main

import "fmt"

func main() {
	isLoggedIn := false
	var userStatus string

	if isLoggedIn {
		userStatus = "Logged In"
		fmt.Println(userStatus)

	} else {
		userStatus = "Not Logged In"
		fmt.Println(userStatus)

	}

	marks := 90

	switch marks {
	case 90:
		fmt.Println("You got A+")
	case 80:
		fmt.Println("You got A")
	case 70:
		fmt.Println("You got B")
	default:
		fmt.Println("You got C")
	}

	// type switch₹
	whoIAm := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am a Integer")
		case string:
			fmt.Println("I am a String")
		case bool:
			fmt.Println("I am a Boolean")
		default:
			fmt.Println("I am of unknown type")

		}
	}

	whoIAm(67)

	// Array

	var myArray [5]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	fmt.Println(myArray)

	fmt.Println(len(myArray))

	// Declare Array in Single line
	myArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray2)

	//  2D Array
	myArray3 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(myArray3)

	// Slices in Go
	// Slices are Dynamic Array
	// In Arr we have fixed size, in Slices it can be expanded and shrink dynamically
	// Slices Mostly used in GO
	// Slices have methods
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	// Append to Slice
	mySlice = append(mySlice, 6)
	fmt.Println(mySlice)

	// // Copy Slice
	// mySlice2 := make([]int, len(mySlice))
	// copy(mySlice2, mySlice)
	// fmt.Println(mySlice2)

	// // Delete from Slice
	// mySlice = append(mySlice[:2], mySlice[3:]...)
	// fmt.Println(mySlice)

	var nums []int

	fmt.Println(nums == nil)

}
