package pointers

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// This is a function named Struct that takes a string parameter name.
// Inside the function, a new bill struct instance b is created and initialized with the provided name, an empty items map, and a tip amount of 0.
// The initialized bill instance b is then returned from the function.

func Struct(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
