package main

import "fmt"

//Branching with if and else in Go is straight-forward.

func main() {
	//Here is a basic example.
	if 9%2 == 0 {
		fmt.Println("9 is even.")
	} else {
		fmt.Println("9 is odd.")
	}

	//You can have an if statement without an else.
	if 16%4 == 0 {
		fmt.Println("16 is divisible by 4.")
	}

	//A statement can precede conditionals;
	//any variables declared in this statement are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative.")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, "has multiple digits.")
	}
}

//Note that you don’t need parentheses around conditions in Go, but that the braces are required.
//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions. (see https://en.wikipedia.org/wiki/%3F: about ternary if)
