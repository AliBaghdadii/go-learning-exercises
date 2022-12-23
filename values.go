package main

import "fmt"

//Declare the values: int, string, boolean, float

func main() {
	//int values
	firstInt := 42
	secondInt := 13
	
	//string values
	firstString := "Hi"
	secondString := "Bye"
	
	//boolean values, we have only two kinds of this values: true, values
	trueBoolean := true
	falseBoolean := false
	
	//float values
	firstFloat := 0.001
	secondFloat := 987.5
	
	fmt.Println(trueBoolean && falseBoolean)
    fmt.Println(trueBoolean || falseBoolean)
    fmt.Println(!trueBoolean)
    
    fmt.Println(firstString + "and" + secondString)
    fmt.Println("42 + 13 =", firstInt + secondInt)
    fmt.Println("987.5 / 0.001 = ", secondFloat / firstFloat)
	}
