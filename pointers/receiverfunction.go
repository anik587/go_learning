package pointers

import "fmt"

type bills struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func NewBill(name string) bills {
	b := bills{
		name:  name,
		items: map[string]float64{"pie": 5.49, "cake": 3.49},
		tip:   0,
	}
	return b
}

// format the bill
func (b bills) Format() string {
	fs := b.name + " Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%0.2v\n", k+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}
