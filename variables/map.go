package variables

import "fmt"

func Map() {
	menu := map[string]float64{
		"nana":   4.99,
		"valir":  7.99,
		"vale":   6.99,
		"vexana": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["vale"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "balmond",
		984759373: "cici",
		845775485: "thamuz",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "jonson"
	fmt.Println(phonebook)

	phonebook[647583927] = "tigrel"
	fmt.Println(phonebook)
}
