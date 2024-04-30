package variables

import "fmt"

func Print() {
	age := 35
	name := "shaun"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")

	fmt.Printf("For Integer %d", age)
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string), %_ = format specifier
	fmt.Printf("my age is %v and my name is %v ==> %%v = value in default format\n", age, name)
	fmt.Printf("my age is %q and my name is %q ==> %%q = quotes \n", age, name)
	fmt.Printf("age is of type %T ==> %%T is the type\n", age)
	fmt.Printf("you scored %f points! ==> %%f = float format \n", 225.55)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	var str = fmt.Sprintf("my age is %v and my name is %v ==> Sprintf (save formatted strings) \n", age, name)
	fmt.Println("the saved string is:", str)

	// see more format specifiers here:
	// https://golang.org/pkg/fmt/
}
