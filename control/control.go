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
	switch day {
	case "Monday":
		fmt.Println("It's a boring Monday")
	case "Tuesday":
		fmt.Println("We are making progress, it's Tuesday")
	default:
		fmt.Println("It's another dat, hopefully weekend")
	}
}
