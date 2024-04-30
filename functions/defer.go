package functions

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

func DeferFunc() {

	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	defer fmt.Println("world")
	fmt.Println("hello")
	loopDefer()
}

func loopDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
