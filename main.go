package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang_ninja/concurrency"
	"golang_ninja/functions"
	"golang_ninja/interfaces"
	"golang_ninja/iterations"
	"golang_ninja/pointers"
	"golang_ninja/variables"
)

var score = 99.5

func printPromt(str string) {
	fmt.Println(str)
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
		//input = strconv.Itoa(rand.Intn(24) + 1)
	}
	fmt.Println(input)
	input = strings.TrimSpace(input)
	return input
}

func getUserChoice(input string) int {
	// Convert input to integer
	choice, err := strconv.Atoi(input) // Remove newline character

	if err != nil {
		fmt.Println("Invalid choice. Please enter a number.")
		return 0
	}
	return choice
}

func main() {

	prompt := []string{"Menu:", "1. Number", "2. Print", "3. Array", "4. String", "5. Loop", "6. Condition",
		"7. Referance function", "8. Return multiple function", "9. Package Scope", "10. Map", "11. Pass By Value", "12. Pointer", "13. Struct", "14. Receiver Function", "15. Receiver Pointers", "16. Save Files", "17. Interface", "18. Type Assertions", "19. Type Switch", "20. Goroutine", "21. defer function", "22. Waiting Group", "23. Mutex", "24. Race Condition", "25. Channel & Deadlock", "30. Exit", "Please enter your choice (number like 12): "}

	for _, v := range prompt {
		printPromt(v)
	}

	input := getUserInput()
	choice := getUserChoice(input)

	switch choice {
	case 1:
		//1. Number
		variables.Number()

	case 2:
		//2. Print
		variables.Print()

	case 3:
		//3. Array
		variables.Array()

	case 4:
		//4. String
		variables.String()

	case 5:
		//5. Loop
		iterations.Loop()

	case 6:
		//6. Condition
		iterations.Condition()

	case 7:
		//7. Referance function
		functions.Ref_function()

	case 8:
		//8. Return multiple function
		functions.Return_muntiple_function()

	case 9:
		//9. Package Scope
		sayHello("mario")
		showScore()

		for _, v := range points {
			fmt.Println(v)
		}

	case 10:
		//10. Map
		variables.Map()

	case 11:
		//11. Pass By Value
		pointers.Passbyvalue()

	case 12:
		//12. Pointer
		pointers.Pointer()

	case 13:
		//13. Struct
		fmt.Println(pointers.Struct("Leomord"))

	case 14:
		//14. Receiver Function
		newBill := pointers.NewBill("Clint's bill")
		fmt.Println(newBill.Format())

	case 15:
		//15. Receiver Pointers
		newBillss := pointers.NewBillss("anik'ss bill")
		newBillss.AddItem("onion soup", 4.50)
		newBillss.AddItem("veg pie", 8.95)
		newBillss.AddItem("toffee pudding", 4.95)
		newBillss.AddItem("coffee", 3.25)

		newBillss.UpdateTip(10)

		fmt.Println(newBillss.Formats())

	case 16:
		//16. Save Files
		saveBill := pointers.NewBillss("Saving Bills")
		saveBill.AddItem("Rice", 90)
		saveBill.AddItem("Veges", 150)

		saveBill.UpdateTip(20)
		saveBill.Save()

		return

	case 17:
		//17. Interface
		shapes := []interfaces.Shape{
			interfaces.Square{Length: 15.2},
			interfaces.Circle{Radius: 7.5},
			interfaces.Circle{Radius: 12.3},
			interfaces.Square{Length: 4.9},
		}
		fmt.Printf("%T\n", shapes)
		fmt.Printf("%v\n", shapes)
		for _, v := range shapes {
			interfaces.PrintShapeInfo(v)
			fmt.Printf("%T", v)
			fmt.Println("---")
		}

	case 18:
		//18. Type Assertions
		interfaces.Typeassertions()

	case 19:
		//19. Type Switch
		interfaces.Typeswitch("string")
		interfaces.Typeswitch(55)
		interfaces.Typeswitch(true)

	case 20:
		//20. Goroutine
		concurrency.Goroutine()

	case 21:
		//21. defer function
		functions.DeferFunc()

	case 22:
		//22. Waiting Group
		concurrency.Waitgroup()

	case 23:
		//23. Mutex
		concurrency.Mutex()

	case 24:
		//24. Race Condition
		concurrency.RaceCondition()
	case 25:
		//25. Channel & Deadlock
		concurrency.Channel()

	case 30:

		fmt.Println("Exiting...")
		return

	default:

		fmt.Println("Not Listed")
		return

	}

}
