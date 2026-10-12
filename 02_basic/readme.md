# Switch Statement & Arrays

## Switch

- In go we do not need to mention `break` keyword
- Go Takecare it automatically

```go
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

```

## Type Switch
- widely used in go 
- You can use switch statement based on type

```go
func main() {

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Printf("I is an Integer & value is %d\n", i)
		case string:
			fmt.Printf("I is a String & value is %s\n", i)
		case bool:
			fmt.Printf("I is Boolean & value is %t\n", i)
		}

	}

	whoAmI("Ankush")
	whoAmI(19)
	whoAmI(true)
}

```

## Arrays
- Array is a numbered sequence of specific length
- we can declare array by specificing the value
- If we do no specify the value, then all other element are treated as `0`, ` '' ` empty string or  `false` value based on the type

- `len` is a method given by go to calculate the length of an array

```go
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	fmt.Println(len(myArray))

	// 2D Array
	myarray2 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(myarray2)
```


