package functions

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func printColor(c string) {
	fmt.Printf("Favourite color is %v \n", c)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func cycleColors(c []string, f func(string)) {
	for _, v := range c {
		f(v)
	}
}

func tryRef(c []string, f func(string)) float64 {
	for _, v := range c {
		f(v)
	}
	return 9.99999
}
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// In Go, fields, variables, types, and functions are exported (i.e., visible outside the package) if their names start with an uppercase letter.
// If a field starts with a lowercase letter, it's unexported and can only be accessed within the same package. Exporting a field by renaming it to start with an uppercase letter allows it to be accessed from outside the package.

func Ref_function() {
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)

	cycleColors([]string{"black", "green", "blue", "yellow"}, printColor)
	tryRef([]string{"peach"}, printColor)
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}
