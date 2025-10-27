package main

import "fmt"

var (
	number1 float64
	number2 float64

	operation string
)

func askOperation() {
	for {
		fmt.Print("What operation you want to do? ")
		fmt.Scanln(&operation)

		if (operation == "+") || (operation == "-") ||
			(operation == "*") || (operation == "/") {
			break
		}

		fmt.Println("These are the only accepted operations: +, -, *, /")
	}

	fmt.Println()
}

func askNumbers() {
	fmt.Print("Input the first number: ")
	fmt.Scanln(&number1)

	for {
		fmt.Print("Input the second number: ")
		fmt.Scanln(&number2)

		if (operation == "/") && (number2 == 0) {
			fmt.Println("The second number can't be 0 if you are trying to do a division")
		} else {
			break
		}
	}

	fmt.Println()
}

func doOperation(internalOperation string, internalNumber1 float64, internalNumber2 float64) float64 {
	var result float64

	switch internalOperation {
	case "+":
		result = internalNumber1 + internalNumber2
	case "-":
		result = internalNumber1 - internalNumber2
	case "*":
		result = internalNumber1 * internalNumber2
	case "/":
		result = internalNumber1 / internalNumber2
	}

	return result
}

func main() {
	fmt.Println("===========================")
	fmt.Println("GO CALCULATOR")
	fmt.Println("===========================")

	askOperation()

	askNumbers()

	result := doOperation(operation, number1, number2)

	fmt.Printf("This is the result: %v %v %v = %v \n", number1, operation, number2, result)
}
