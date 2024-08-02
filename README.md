# Go Programming Course

Welcome to the Go Programming Course! This guide will walk you through the fundamental concepts of Go, from writing a basic "Hello, World!" program to working with packages and structs.

## Prerequisites

- Go installed on your computer
- A text editor or IDE (such as VSCode, GoLand, or Sublime Text)

## Directory Structure

```
go-project/
├── main/
│ └── main.go
├── variables/
│ └── variables.go
├── control_structures/
│ └── control_structures.go
├── functions/
│ └── functions.go
└── structs/
└── structs.go
```

## 1. Getting Started

### 1. Hello, World!

1. **Create a directory:**
   ```sh
   mkdir go-hello-world
   cd go-hello-world
   ```

### 2. Create main.go:

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 3. Run the program

```sh
go run main.go

```

## 2. Variables and Data Types

Create a directory for this example:

```sh
mkdir variables
cd variables
Create variables.go:
```

```go
// variables.go
package main

import "fmt"

func main() {
	age := 33
	fullName := "Ivan Rodriguez"
	isMarried := true

	fmt.Printf("Hello %s, you are %v years old and is %v that you are married", fullName, age, isMarried)

```

Run your program:

```sh
go run variables.go
```

## 3. Control Structures

Create a directory for this example:

```sh
mkdir control
cd control
touch control.go:
```

```go
// control.go
package main

import "fmt"

func main() {
	age := 19

	// If-Else
	if age > 18 {
		fmt.Println("You are old enough for an alcoholic drink")
	} else {
		fmt.Println("Sorry, you are too young")
	}

	// For-loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("I am iteration # %v", i)
		fmt.Println()
	}

	day := "Tuesday"

  // Switch statements
	switch day {
	case "Monday":
		fmt.Println("It's a boring Monday")
	case "Tuesday":
		fmt.Println("We are making progress, it's Tuesday")
	default:
		fmt.Println("It's another dat, hopefully weekend")
	}
}
```

Run your program:

```sh
go run control_structures.go
```

## 4. Functions

Create a directory for this example:

```sh
mkdir functions
cd functions
Create functions.go:
```

```go
// functions.go
package main

import "fmt"

func add(a int, b int) int {
return a + b
}

func swap(x, y string) (string, string) {
return y, x
}

func main() {
sum := add(5, 3)
fmt.Println("Sum:", sum)

    a, b := swap("hello", "world")
    fmt.Println("Swapped:", a, b)

}
```

Run your program:

```sh
go run functions.go
```

## 5. Structs and Methods

Create a directory for this example:

```sh
mkdir structs
cd structs
Create structs.go:
```

```go
// structs.go
package main

import "fmt"

type Person struct {
firstName string
lastName string
age int
}

func (p Person) fullName() string {
return p.firstName + " " + p.lastName
}

func main() {
person := Person{firstName: "John", lastName: "Doe", age: 30}

    fmt.Println("First Name:", person.firstName)
    fmt.Println("Last Name:", person.lastName)
    fmt.Println("Age:", person.age)
    fmt.Println("Full Name:", person.fullName())

}
```

Run your program:

```sh
go run structs.go
```

## 6. Packages

Create a directory for your package:

```sh
mkdir mypackage
cd mypackage
Create mypackage.go:
```

```go
// mypackage/mypackage.go
package mypackage

import "fmt"

func Hello() {
fmt.Println("Hello from mypackage!")
}
```

Create main.go to use your package:

```go
// main.go
package main

import "mypackage"

func main() {
mypackage.Hello()
}
```

Run your program:

```sh
go run main.go
```

Happy coding!
