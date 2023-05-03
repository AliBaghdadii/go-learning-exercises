package main

import (
	"errors"
	"fmt"
)

/*In Go it’s idiomatic to communicate errors via an explicit, 
separate return value. 
This contrasts with the exceptions used in languages like Java and Ruby 
and the overloaded single result / error value sometimes used in C. 
Go’s approach makes it easy to see which functions return errors 
and to handle them using the same language constructs employed for any other, non-error tasks.*/

//By convention, errors are the last return value and have type error, a built-in interface.

func f1(arg int) (int, error) {

	//errors.New constructs a basic error value with the given error message.

	if arg == 13 {
		return -1, errors.New("Can't work with 13.")
	}
	
	//A nil value in the error position indicates that there was no error.

	return arg + 13, nil
}

/*It’s possible to use custom types as errors by implementing the Error() method on them. 
Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.*/

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 13 {
		return -1, &argError{arg, "Can't work with it."}
	}
	return arg + 13, nil
}

func main() {

	/*The two loops below test out each of our error-returning functions. 
	Note that the use of an inline error check on the if line is a common idiom in Go code.*/

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e!= nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	/*If you want to programmatically use the data in a custom error, 
	you’ll need to get the error as an instance of the custom error type via type assertion.*/

	_, e := f2(13)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}