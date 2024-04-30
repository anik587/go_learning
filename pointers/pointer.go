package pointers

import "fmt"

func updateData(n *string) {
	*n = "wedge"
}

func Pointer() {
	name := "tifa"
	//updateData(&name)
	fmt.Println("memory address of name is:", &name)
	m := &name
	fmt.Println("memory address before", m)
	fmt.Println("value at memory address before:", *m)

	updateData(m)

	fmt.Println("memory address after", m)
	fmt.Println("value at memory address after:", *m)
}
