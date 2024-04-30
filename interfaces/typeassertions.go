package interfaces

import "fmt"

func Typeassertions() {
	var i interface{} = "hello"

	// Type assertion to extract the underlying string value
	s, ok := i.(string)
	if ok {
		fmt.Println("Value:", s)
	} else {
		fmt.Println("Assertion failed")
	}
	// 	We have an interface variable i holding the value "hello".
	// We use a type assertion to extract the underlying string value from the interface variable i.
	// If the assertion succeeds, we print the extracted string value. Otherwise, we print "Assertion failed".
}
