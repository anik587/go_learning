package variables

import (
	"fmt"
	"sort"
	"strings"
)

func String() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "bye"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println("Ints sorts in increasing order", ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println("SearchInts searches for x in a sorted slice", index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println("SearchStrings searches for x in a sorted slice", sort.SearchStrings(names, "mario"))

	anotherArray := []float64{1.1, 2.2, 3.3}
	fmt.Println(anotherArray)
}
