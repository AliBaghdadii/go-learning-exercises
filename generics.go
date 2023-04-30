package main

import "fmt"

/*Starting with version 1.18, 
Go has added support for generics, also known as type parameters.*/

/*As an example of a generic function, 
MapKeys takes a map of any type and returns a slice of its keys. 
This function has two type parameters - K and V; K has the comparable constraint, 
meaning that we can compare values of this type with the == and != operators. 
This is required for map keys in Go. V has the any constraint, 
meaning that it’s not restricted in any way (any is an alias for interface{}).*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

/*As an example of a generic type, 
List is a singly-linked list with values of any type.*/
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

/*We can define methods on generic types just like we do on regular types, 
but we have to keep the type parameters in place. The type is List[T], not List.*/
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	/*When invoking generic functions, 
	we can often rely on type inference. 
	Note that we don’t have to specify the types for K and V when calling 
	MapKeys - the compiler infers them automatically.*/

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(12)
	lst.Push(18)
	lst.Push(27)

	fmt.Println("list:", lst.GetAll())
}