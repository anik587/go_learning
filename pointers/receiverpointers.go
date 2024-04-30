package pointers

import (
	"fmt"
	"os"
)

type billss struct {
	name  string
	items map[string]float64
	tips  float64
}

func NewBillss(name string) billss {
	b := billss{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}
	return b
}

func (b *billss) AddItem(name string, price float64) {
	b.items[name] = price
}

func (b *billss) UpdateTip(tip float64) {
	(*b).tips = tip
}

// format the bill
func (b *billss) Formats() string {
	fmt.Println(b)
	fs := b.name + "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tips)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tips)

	return fs
}

func (b *billss) Save() {
	data := []byte(b.Formats())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}
