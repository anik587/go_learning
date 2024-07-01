package interfaces

import "fmt"

func Typeswitch(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}

func Fallthrough() {
	i := 2
	a := 5
	b := 4
	switch i {
	case 1:
		fmt.Println("from 1")
		if b == 4 {
			fmt.Println("from 1+4")
			goto case2
		}
	case2:
		fmt.Println("case2")
	case 2:
		fmt.Println("before break from 2")
		if a == 5 {
			fmt.Println("before break")
			break
		}
		fallthrough
	case 3:
		fmt.Println("from 3")
	case 4:
		fmt.Println("from 4")
		fallthrough
	default:
		fmt.Println("from default")

	}
}
