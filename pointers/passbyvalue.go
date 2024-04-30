package pointers

import "fmt"

func updateName(x string) string {
	fmt.Println(x)
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func Passbyvalue() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"

	// updateName(name)
	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	// pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
