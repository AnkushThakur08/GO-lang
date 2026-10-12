package main

import "fmt"

func main() {
	var myName string = "Ankush Thakur"
	var myAge int = 28
	var isMarried bool = false

	var marriedStatus string = "UnMarried"

	if isMarried {
		marriedStatus = "Married"
	}

	// fmt.Println(`My Name is ` + myName + ` and I am ` + fmt.Sprint(myAge) + ` years old &  I am ` + marriedStatus)
	fmt.Printf("My Name is %s, I am %d years old & my Marital status is %s\n", myName, myAge, marriedStatus)

	fmt.Println("Hello World from Go")

	// Shorthand syntax for variable declaration
	myName2 := "Ankush Thakur"
	myAge2 := 28
	isMarried2 := false

	marriedStatus2 := "UnMarried"

	if isMarried2 {
		marriedStatus2 = "Married"
	}

	fmt.Printf("My Name is %s, I am %d years old & my Marital status is %s\n", myName2, myAge2, marriedStatus2)

	// Constants Grouping

	const (
		port     = 5000
		host     = "localhost"
		db       = "mysql"
		password = 123
	)

	fmt.Println(port, host, db, password)

	// Looping in go
	//  We have only for loop in go, which is used like while

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//  we can also use for loop like this
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// we can use range as well, it is used to iterate over a specific ranfe

	for i := range 10 {
		fmt.Println(i)
	}
}
