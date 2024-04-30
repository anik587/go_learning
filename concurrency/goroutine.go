package concurrency

import (
	"fmt"
)

func Goroutine() {
	go greetings("hello")
	greetings("world")

	// greetings("hello") // only hello will be printed
	// go greetings("world") // because main func return before goroutine and we are not waiting

	// go greetings("hello") // nothing will be print
	// go greetings("world") // because main func return before goroutine
}

func greetings(s string) {
	for i := 0; i < 6; i++ {
		// time.Sleep(time.Second) //--> comment and uncomment to change
		fmt.Printf("%v %v \n", i, s)
	}
}
