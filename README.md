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
var a int = 10
var b string = "Hello"
var c bool = true
d := 20.5

    fmt.Println("a =", a)
    fmt.Println("b =", b)
    fmt.Println("c =", c)
    fmt.Println("d =", d)

}
```

Run your program:

```sh
go run variables.go
```

## 3. Control Structures

Create a directory for this example:

```sh
mkdir control_structures
cd control_structures
Create control_structures.go:
```

```go
// control_structures.go
package main

import "fmt"

func main() {
x := 10
if x > 5 {
fmt.Println("x is greater than 5")
} else {
fmt.Println("x is less than or equal to 5")
}

    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    day := "Tuesday"
    switch day {
    case "Monday":
        fmt.Println("It's Monday")
    case "Tuesday":
        fmt.Println("It's Tuesday")
    default:
        fmt.Println("It's some other day")
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
