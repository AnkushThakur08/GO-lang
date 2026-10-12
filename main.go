package main

import "fmt"

func main() {

	// whoAmI := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		fmt.Printf("I is an Integer & value is %d\n", i)
	// 	case string:
	// 		fmt.Printf("I is a String & value is %s\n", i)
	// 	case bool:
	// 		fmt.Printf("I is Boolean & value is %t\n", i)
	// 	}

	// }

	// whoAmI("Ankush")
	// whoAmI(19)
	// whoAmI(true)

	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	fmt.Println(len(myArray))

	myarray2 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(myarray2)
}
