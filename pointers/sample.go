package main

import (
	"fmt"
)

func MultipleOfTen(factor *int) *int {
	v := *factor * 10
	return &v
}

func main() {
	var a int = 10
    var ptr *int
    ptr = &a

	fmt.Println("===== Example 1 =====")
    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", &a)
    fmt.Println("Value of ptr:", ptr)
    fmt.Println("Value pointed to by ptr:", *ptr)

	fmt.Println("===== Example 2 =====")
	fmt.Println("Value of a before:", a)
    *ptr = 20
    fmt.Println("Value of a after:", a)
	
	fmt.Println("===== Example 3 =====")
	var p int = 40
	addr := MultipleOfTen(&p)
	val := *MultipleOfTen(&p)
	fmt.Println("Address:", addr)
	fmt.Println("Return value:", val)
}