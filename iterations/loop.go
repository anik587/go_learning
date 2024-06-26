package iterations

import (
	"fmt"
)

func Loop() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}
	fmt.Println()

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println()

	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		fmt.Printf("the type at pos %v is %T \n", index, val)
		val = "new string"
	}

	for _, val := range names {
		fmt.Println(val, ",")
		val = "new string" // work because of scope variable
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)
}
