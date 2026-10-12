# GO Lang

### How to start

- Go Lang code starts like this
- we have `package main`
- & the `main` function the entry point

```go
package main

func main() {
    //  func main is the entry Point
}
```

### Command to run the file

`go run <fileName>`

### How to Print/Console in go Lang

`fmt.Println("Hello World from Go")`

## Variable & Interpolation in Go Lang

- Int
- String
- bool

| Format Specifier | Data Type | Example          |
| ---------------- | --------- | ---------------- |
| `%s`             | String    | `"Ankush"`       |
| `%d`             | Integer   | `28`             |
| `%t`             | Boolean   | `true` / `false` |
| `%f`             | Float     | `10.5`           |

### Quick Reference

- `%s` → String
- `%d` → Integer
- `%t` → Boolean
- `%f` → Float

```go
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
}
```

## Constant Grouping

- In Go lang, we declare variable using `const` to make it constant
- We can also group multiple const value which is know as `const Grouping`

```go
	const (
		port     = 5000
		host     = "localhost"
		db       = "mysql"
		password = 123
	)

	fmt.Println(port, host, db, password)

```


## Looping
- we have for loop in go
- There are 2 different syntax for for loop in go

### For while loop
```go
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
```

### For
```go
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
```