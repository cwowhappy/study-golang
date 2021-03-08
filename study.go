package main

import "fmt"

type A struct {
	value int
}

func main()  {
	func1 := func(a A) {
		fmt.Printf("[func] Point=%p \n", &a)
	}

	var a = A{10}
	fmt.Printf("[main] Point=%p \n", &a)
	func1(a)
}
