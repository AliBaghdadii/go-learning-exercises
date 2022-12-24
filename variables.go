package main

import "fmt"

func main() {
	//var declares 1 or more variables.
	var a = "first one"
    fmt.Println(a)
    
	//You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)
    
	//Go will infer the type of initialized variables.
    var a_Boolean = true
    fmt.Println(!a_Boolean)
    
	//Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var d int
    fmt.Println(d)
    
	//The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
    e := "apple"
    fmt.Println(e)
	}
